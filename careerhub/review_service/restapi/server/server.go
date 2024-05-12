package server

import (
	"context"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/domain/company"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi/repo"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi/restapi_grpc"
)

type ReviewReaderGrpcServer struct {
	companyRepo *repo.CompanyRepo
	reviewRepo  *repo.ReviewRepo
	restapi_grpc.UnimplementedReviewReaderGrpcServer
}

func NewReviewReaderGrpcServer(companyRepo *repo.CompanyRepo, reviewRepo *repo.ReviewRepo) *ReviewReaderGrpcServer {
	return &ReviewReaderGrpcServer{
		companyRepo: companyRepo,
		reviewRepo:  reviewRepo,
	}
}

func (s *ReviewReaderGrpcServer) GetCompanyScores(ctx context.Context, in *restapi_grpc.GetCompanyScoresRequest) (*restapi_grpc.GetCompanyScoresResponse, error) {
	companyNamesMap := make(map[string]string)
	defaultCompanyNames := make([]string, 0, len(in.CompanyNames))
	for _, companyName := range in.CompanyNames {
		defaultName := company.RefineNameForSearch(companyName)
		companyNamesMap[defaultName] = companyName
		defaultCompanyNames = append(defaultCompanyNames, defaultName)
	}

	companies, err := s.companyRepo.GetCompanies(ctx, in.Site, defaultCompanyNames)
	if err != nil {
		return nil, err
	}

	companyScores := make(map[string]int32)
	for _, c := range companies {
		for _, score := range c.ReviewSites {
			if score.Site == in.Site {
				companyScores[companyNamesMap[c.DefaultName]] = score.AvgScore
			}
		}
	}

	return &restapi_grpc.GetCompanyScoresResponse{
		CompanyScores: companyScores,
	}, nil
}

func (s *ReviewReaderGrpcServer) GetCompanyReviews(ctx context.Context, in *restapi_grpc.GetCompanyReviewsRequest) (*restapi_grpc.GetCompanyReviewsResponse, error) {
	if in.Limit <= 0 {
		in.Limit = 100
	}

	reviews, err := s.reviewRepo.GetReviews(ctx, in.Site, in.CompanyName, in.Offset, in.Limit)
	if err != nil {
		return nil, err
	}

	grpcReviews := make([]*restapi_grpc.Review, 0, len(reviews))
	for _, r := range reviews {
		grpcReviews = append(grpcReviews, &restapi_grpc.Review{
			Score:            r.Score,
			Summary:          r.Summary,
			EmploymentStatus: r.EmploymentStatus,
			ReviewUserId:     r.ReviewUserId,
			JobType:          r.JobType,
			UnixMilli:        r.Date.UnixMilli(),
		})
	}

	return &restapi_grpc.GetCompanyReviewsResponse{
		Reviews: grpcReviews,
	}, nil
}
