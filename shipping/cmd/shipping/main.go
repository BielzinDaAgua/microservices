package main

import (
	"log"
	"net"

	"github.com/bielzindaagua/microservices/payment/internal/adapters/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// Register your shipping service here

	log.Println("Shipping service is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
