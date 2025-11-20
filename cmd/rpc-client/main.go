package main

import (
	"context"
	"log"
	"time"

	pb "simulation/pkg/protocol"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Could not close connection: %v", err)
		}
	}(conn)

	client := pb.NewSimulationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Sending request...")
	res, err := client.GetSalesMetrics(ctx, &pb.SalesRequest{
		Region: "Europe",
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Response: %s (Count: %d)", res.RegionProcessed, res.TotalSalesCount)
}
