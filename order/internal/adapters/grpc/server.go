package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/bielzindaagua/microservices-proto/golang/order"
	"github.com/bielzindaagua/microservices/order/config"
	"github.com/bielzindaagua/microservices/order/internal/application/core/domain"
	"github.com/bielzindaagua/microservices/order/internal/ports"
	grpcstd "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type InventoryChecker interface {
	ItemExists(itemID string) (bool, error)
}

type Adapter struct {
	api           ports.APIPort
	inventoryRepo InventoryChecker
	port          int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, inventoryRepo InventoryChecker, port int) *Adapter {
	return &Adapter{api: api, inventoryRepo: inventoryRepo, port: port}
}

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		exists, err := a.inventoryRepo.ItemExists(orderItem.ProductCode)
		if err != nil {
			return nil, fmt.Errorf("failed to check inventory: %w", err)
		}
		if !exists {
			return nil, fmt.Errorf("item %s does not exist in inventory", orderItem.ProductCode)
		}
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(int64(request.CostumerId), orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{OrderId: int32(result.ID)}, nil
}

func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	grpcServer := grpcstd.NewServer()
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve grpc on port %d", a.port)
	}
}
