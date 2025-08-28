package main

import (
	"log"

	"github.com/bielzindaagua/microservices/order/config"
	"github.com/bielzindaagua/microservices/order/internal/adapters/db"
	grpctransport "github.com/bielzindaagua/microservices/order/internal/adapters/grpc"
	paymentadapter "github.com/bielzindaagua/microservices/order/internal/adapters/payment"
	"github.com/bielzindaagua/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database. error: %v", err)
	}

	paymentAdapter, err := paymentadapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("failed to initialize payment stub. error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpctransport.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()

}
