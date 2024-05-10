package server

import (
	"context"
	"slices"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi/repo"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi/restapi_grpc"
)

type ReviewReaderGrpcServer struct {
	companyRepo *repo.CompanyRepo
	restapi_grpc.UnimplementedReviewReaderGrpcServer
}

func NewReviewReaderGrpcServer(companyRepo *repo.CompanyRepo) *ReviewReaderGrpcServer {
	return &ReviewReaderGrpcServer{
		companyRepo: companyRepo,
	}
}

func (s *ReviewReaderGrpcServer) GetCompanyScores(ctx context.Context, in *restapi_grpc.GetCompanyScoresRequest) (*restapi_grpc.GetCompanyScoresResponse, error) {
	companies, err := s.companyRepo.GetCompanies(ctx, in.Site, in.CompanyNames)
	if err != nil {
		return nil, err
	}

	companyScores := make(map[string]int32)
	for _, companyName := range in.CompanyNames {
		for _, company := range companies {
			if slices.Contains(company.OtherNames, companyName) {
				for _, reviewSite := range company.ReviewSites {
					if reviewSite.Site == in.Site {
						companyScores[companyName] = reviewSite.AvgScore
						break
					}
				}
			}
		}
	}

	return &restapi_grpc.GetCompanyScoresResponse{
		CompanyScores: companyScores,
	}, nil
}
