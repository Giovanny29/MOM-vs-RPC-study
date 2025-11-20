package main

import (
	"context"
	"log"
	"net"

	pb "simulation/pkg/protocol"

	"google.golang.org/grpc"
)

type LoadBalancerServer struct {
	pb.UnimplementedSimulationServiceServer
}

func (s *LoadBalancerServer) GetSalesMetrics(ctx context.Context, req *pb.SalesRequest) (*pb.SalesResponse, error) {
	log.Printf("LB Received Request: Region=%s, Category=%s", req.Region, req.Category)

	// TODO: Forward to a worker

	// For now, return a dummy response to test the connection
	return &pb.SalesResponse{
		TotalSalesCount: 0,
		TotalRevenue:    0,
		RegionProcessed: "LoadBalancer-Dummy",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSimulationServiceServer(grpcServer, &LoadBalancerServer{})

	log.Println("RPC Load Balancer running on :50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
