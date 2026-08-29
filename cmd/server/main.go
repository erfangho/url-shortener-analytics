package main

import (
	"log"
	"net"

	grpcserver "github.com/erfangho/url-shortener-analytics/internal/grpc"
	pb "github.com/erfangho/url-shortener-analytics/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	analyticsServer := grpcserver.NewAnalyticServer()
	pb.RegisterAnalyticsServiceServer(server, analyticsServer)

	log.Println("gRPC server running on :50051")

	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
