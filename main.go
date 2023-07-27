package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"

	"google.golang.org/grpc"

        pb "github.com/amirh/stocks-broker/proto"
)

// Server implements the StocksBroker service
type Server struct{
  pb.UnimplementedStocksBrokerServer
}

func (s *Server) SetLimitOrder(ctx context.Context, order *pb.LimitOrder) (*pb.OrderStatus, error) {
	fmt.Printf("Received Limit Order: %v\n", order)

	// Simulate the price checking with random values (replace this with real API calls)
	currentPrice := rand.Float64() * 100 // Random price between 0 and 100

	if order.Price >= currentPrice {
		return &pb.OrderStatus{
			OrderId:  "order123", // Replace this with an actual order ID generation logic
			Executed: true,
		}, nil
	}

	return &pb.OrderStatus{
		Executed: false,
	}, nil
}

func main() {
	// Start the gRPC server
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Errorf("failed to listen: %v", err))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStocksBrokerServer(grpcServer, &Server{})

	fmt.Printf("Starting Stocks Broker server on port %d...\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Errorf("failed to serve: %v", err))
	}
}
