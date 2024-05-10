package server

import (
	"context"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/repo"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *ReviewGrpcServer) GetCrawlingTasks(ctx context.Context, _ *emptypb.Empty) (*crawler_grpc.GetCrawlingTasksResponse, error) {
	companies, err := s.companyRepo.GetCrawlingTasks(ctx)

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
	_, err := s.companyRepo.SetScoreNPage(context.Background(), in.DefaultName, in.TotalPageCount, in.AvgScore)
	if err != nil {
		return nil, err
	}
	return &crawler_grpc.SetScoreNPageResponse{}, nil
}
