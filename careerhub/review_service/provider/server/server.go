package server

import (
	"context"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/repo"
	"github.com/jae2274/careerhub-review-service/common/domain/company"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProviderGrpcServer struct {
	companyRepo *repo.CompanyRepo
	provider_grpc.UnimplementedCrawlingTaskGrpcServer
}

func NewProviderGrpcServer(companyRepo *repo.CompanyRepo) *ProviderGrpcServer {
	return &ProviderGrpcServer{
		companyRepo: companyRepo,
	}
}

func (p *ProviderGrpcServer) AddCrawlingTask(ctx context.Context, in *provider_grpc.AddCrawlingTaskRequest) (*provider_grpc.AddCrawlingTaskResponse, error) {
	err := p.companyRepo.AddCompany(ctx, &company.Company{
		Name: "testCompany",
	})

	var status string = "created"
	if err != nil {
		if !mongo.IsDuplicateKeyError(err) {
			return nil, err
		}

		status = "duplicated"
	}

	return &provider_grpc.AddCrawlingTaskResponse{
		Status: status,
	}, nil

}
