// main_test.go

package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"math/rand"

        pb "github.com/amirh/stocks-broker/proto"
)

func TestSetLimitOrder(t *testing.T) {
	// Start the gRPC server for testing
	go main()

	// Create a gRPC client connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewStocksBrokerClient(conn)

	// Use a fixed seed for the random number generator in the test
	rand.Seed(40) // Set the seed to any value you want for deterministic tests

	// Test case 1: Simulated order execution (limit price higher than current price)
	order1 := &pb.LimitOrder{
		Symbol:   "AAPL",
		Price:    30.0,
		Quantity: 10,
	}

	status1, err := client.SetLimitOrder(context.Background(), order1)
	assert.NoError(t, err)
	assert.NotNil(t, status1)
	assert.True(t, status1.Executed)

	// Test case 2: Simulated order not executed (limit price lower than current price)
	order2 := &pb.LimitOrder{
		Symbol:   "GOOG",
		Price:    30.0,
		Quantity: 5,
	}

	status2, err := client.SetLimitOrder(context.Background(), order2)
	assert.NoError(t, err)
	assert.NotNil(t, status2)
	assert.False(t, status2.Executed)
}
