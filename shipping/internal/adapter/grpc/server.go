package grpc

import (
	"context"
	"net"

	"github.com/bielzindaagua/microservices-proto/golang/shipping"
	"github.com/bielzindaagua/microservices/payment/internal/adapters/grpc"
)

type Server struct {
	proto.UnimplementedShippingServiceServer
	appService application.ShippingService
}

func NewServer(appService application.ShippingService) *Server {
	return &Server{
		appService: appService,
	}
}

func (s *Server) CalculateDeliveryTime(ctx context.Context, req *proto.CalculateDeliveryTimeRequest) (*proto.CalculateDeliveryTimeResponse, error) {
	deliveryTime, err := s.appService.CalculateDeliveryTime(req.Items, req.PurchaseId)
	if err != nil {
		return nil, err
	}
	return &proto.CalculateDeliveryTimeResponse{DeliveryTime: deliveryTime}, nil
}

func (s *Server) Start(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	proto.RegisterShippingServiceServer(grpcServer, s)
	return grpcServer.Serve(listener)
}
