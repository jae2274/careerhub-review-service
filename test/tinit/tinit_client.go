package tinit

import (
	"testing"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/provider_grpc"
)

func InitCrawlingTaskGrpcClient(t *testing.T) crawler_grpc.ReviewGrpcClient {
	envVars := InitEnvVars(t)
	conn := InitGrpcClient(t, envVars.CrawlerGrpcPort)

	return crawler_grpc.NewReviewGrpcClient(conn)
}

func InitProviderGrpcClient(t *testing.T) provider_grpc.CrawlingTaskGrpcClient {
	envVars := InitEnvVars(t)
	conn := InitGrpcClient(t, envVars.ProviderGrpcPort)

	return provider_grpc.NewCrawlingTaskGrpcClient(conn)
}
