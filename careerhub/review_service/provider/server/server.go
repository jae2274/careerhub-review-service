package server

import "github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/provider_grpc"

type ProviderGrpcServer struct {
	provider_grpc.UnimplementedCrawlingTaskGrpcServer
}

func NewProviderGrpcServer() *ProviderGrpcServer {
	return &ProviderGrpcServer{}
}
