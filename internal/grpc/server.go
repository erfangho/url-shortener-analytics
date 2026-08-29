package grpc

import (
	"context"
	"log"

	pb "github.com/erfangho/url-shortener-analytics/proto"
)

type AnalyticServer struct {
	pb.UnimplementedAnalyticsServiceServer
}

func NewAnalyticServer() *AnalyticServer {
	return &AnalyticServer{}
}

func (s *AnalyticServer) RecordClick(ctx context.Context, event *pb.ClickEvent) (*pb.RecordResponse, error) {
	log.Printf("Received click: URL_ID:%d", event.UrlId)

	return &pb.RecordResponse{
		Success: true,
		Message: "click recorded",
	}, nil
}
