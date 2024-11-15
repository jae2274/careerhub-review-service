package server

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/domain/company"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/domain/review"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/repo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ReviewGrpcServer struct {
	companyRepo *repo.CompanyRepo
	reviewRepo  *repo.ReviewRepo
	crawler_grpc.UnimplementedReviewGrpcServer
}

func NewReviewGrpcServer(
	companyRepo *repo.CompanyRepo,
	reviewRepo *repo.ReviewRepo,
) *ReviewGrpcServer {
	return &ReviewGrpcServer{
		companyRepo: companyRepo,
		reviewRepo:  reviewRepo,
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

func (s *ReviewGrpcServer) SetReviewScore(ctx context.Context, in *crawler_grpc.SetReviewScoreRequest) (*emptypb.Empty, error) {
	err := s.companyRepo.SetReviewScore(context.Background(), company.RefineNameForSearch(in.CompanyName), &company.ReviewSite{
		Site:           in.Site,
		AvgScore:       in.AvgScore,
		ReviewCount:    in.ReviewCount,
		TotalPageCount: in.TotalPageCount,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *ReviewGrpcServer) SetNotExist(ctx context.Context, in *crawler_grpc.SetNotExistRequest) (*emptypb.Empty, error) {
	err := s.companyRepo.SetNotExist(context.Background(), company.RefineNameForSearch(in.CompanyName), in.Site)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *ReviewGrpcServer) GetCrawlingTargets(ctx context.Context, in *crawler_grpc.GetCrawlingTargetsRequest) (*crawler_grpc.GetCrawlingTargetsResponse, error) {
	companies, err := s.companyRepo.GetCrawlingTargets(ctx, in.Site)
	if err != nil {
		return nil, err
	}

	targets := make([]*crawler_grpc.CrawlingTarget, 0)
	for _, c := range companies {
		for _, reviewSite := range c.ReviewSites {
			if reviewSite.Site == in.Site {
				targets = append(targets, &crawler_grpc.CrawlingTarget{
					CompanyName:    c.DefaultName,
					TotalPageCount: reviewSite.TotalPageCount,
				})
				break
			}
		}
	}
	return &crawler_grpc.GetCrawlingTargetsResponse{
		Targets: targets,
	}, nil
}

func (s *ReviewGrpcServer) SaveCompanyReviews(ctx context.Context, in *crawler_grpc.SaveCompanyReviewsRequest) (*crawler_grpc.SaveCompanyReviewsResponse, error) {
	reviews := make([]*review.Review, 0, len(in.Reviews))
	for _, r := range in.Reviews {
		reviews = append(reviews, &review.Review{
			Site:             in.Site,
			CompanyName:      in.CompanyName,
			Score:            r.Score,
			Summary:          r.Summary,
			EmploymentStatus: r.EmploymentStatus,
			ReviewUserId:     r.ReviewUserId,
			JobType:          r.JobType,
			Date:             time.UnixMilli(r.UnixMilli),
		})
	}

	insertedCount, err := s.reviewRepo.InsertReviews(ctx, reviews)
	return &crawler_grpc.SaveCompanyReviewsResponse{InsertedCount: int32(insertedCount)}, err
}

func (s *ReviewGrpcServer) FinishCrawlingTask(ctx context.Context, in *crawler_grpc.FinishCrawlingTaskRequest) (*emptypb.Empty, error) {
	err := s.companyRepo.FinishCrawlingTask(ctx, in.CompanyName, in.Site)

	return &emptypb.Empty{}, err
}
