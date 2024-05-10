package server

import (
	"context"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/repo"
	"github.com/jae2274/careerhub-review-service/common/domain/company"
)

type CrawlingTaskGrpcServer struct {
	companyRepo *repo.CompanyRepo
	provider_grpc.UnimplementedCrawlingTaskGrpcServer
}

func NewCrawlingTaskGrpcServer(companyRepo *repo.CompanyRepo) *CrawlingTaskGrpcServer {
	return &CrawlingTaskGrpcServer{
		companyRepo: companyRepo,
	}
}

const (
	CrawlingTaskCreated     = "created"
	CrawlingTaskDuplicated  = "duplicated"
	CrawlingTaskNotModified = "not_modified"
)

func (p *CrawlingTaskGrpcServer) AddCrawlingTask(ctx context.Context, in *provider_grpc.AddCrawlingTaskRequest) (*provider_grpc.AddCrawlingTaskResponse, error) {
	refinedName := company.RefineNameForSearch(in.CompanyName)
	_, isExisted, err := p.companyRepo.FindCompany(ctx, refinedName, in.CompanyName)
	if err != nil {
		return nil, err
	}

	res := &provider_grpc.AddCrawlingTaskResponse{}
	if isExisted {
		res.Status = CrawlingTaskDuplicated
		return res, nil
	}

	result, err := p.companyRepo.Save(ctx, refinedName, in.CompanyName)

	if err != nil {
		return nil, err
	}

	if result.UpsertedCount > 0 {
		res.Status = CrawlingTaskCreated
	} else if result.ModifiedCount > 0 {
		res.Status = CrawlingTaskNotModified
	}

	return res, nil
}
