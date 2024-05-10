package server

import (
	"context"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/domain/company"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/repo"
)

type ReviewGrpcServer struct {
	companyRepo *repo.CompanyRepo
	crawler_grpc.UnimplementedReviewGrpcServer
}

func NewReviewGrpcServer(
	companyRepo *repo.CompanyRepo,
) *ReviewGrpcServer {
	return &ReviewGrpcServer{
		companyRepo: companyRepo,
	}
}

func (s *ReviewGrpcServer) GetCrawlingTasks(ctx context.Context, in *crawler_grpc.GetCrawlingTasksRequest) (*crawler_grpc.GetCrawlingTasksResponse, error) {
	companies, err := s.companyRepo.GetCrawlingTasks(ctx, in.Site)

	if err != nil {
		return nil, err
	}

	companyNames := make([]string, 0, len(companies))
	for _, company := range companies {
		companyNames = append(companyNames, company.DefaultName)
	}

	return &crawler_grpc.GetCrawlingTasksResponse{
		CompanyNames: companyNames,
	}, nil
}

func (s *ReviewGrpcServer) SetScoreNPage(ctx context.Context, in *crawler_grpc.SetScoreNPageRequest) (*crawler_grpc.SetScoreNPageResponse, error) {
	err := s.companyRepo.SetScoreNPage(context.Background(), in.CompanyName, &company.ReviewSite{
		Site:                in.Site,
		Status:              company.Exist,
		AvgScore:            in.AvgScore,
		CurrentCrawlingPage: in.TotalPageCount,
	})
	if err != nil {
		return nil, err
	}
	return &crawler_grpc.SetScoreNPageResponse{}, nil
}

func (s *ReviewGrpcServer) SetNotExist(ctx context.Context, in *crawler_grpc.SetNotExistRequest) (*crawler_grpc.SetNotExistResponse, error) {
	err := s.companyRepo.SetNotExist(context.Background(), in.CompanyName, in.Site)
	if err != nil {
		return nil, err
	}
	return &crawler_grpc.SetNotExistResponse{}, nil
}
