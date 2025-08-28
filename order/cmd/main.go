package main

import (
	"log"

	"github.com/bielzindaagua/microservices/order/config"
	db "github.com/bielzindaagua/microservices/order/internal/adapters/db"
	grpc "github.com/bielzindaagua/microservices/order/internal/adapters/grpc"
	payment_adapter "github.com/bielzindaagua/microservices/order/internal/adapters/payment"
	"github.com/bielzindaagua/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf(" Failed to connect to database . Error : %v", err)
	}
	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf(" Failed to initialize payment stub . Error : %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
