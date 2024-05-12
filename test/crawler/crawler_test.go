package crawler

import (
	"context"
	"testing"
	"time"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-review-service/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestReviewGrpcClient(t *testing.T) {
	cancelFunc := tinit.RunTestApp(t)
	defer cancelFunc()

	blindSite := "blind"
	t.Run("return empty tasks when nothing saved", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)

		res, err := client.GetCrawlingTasks(ctx, &crawler_grpc.GetCrawlingTasksRequest{Site: blindSite})
		require.NoError(t, err)
		require.Empty(t, res.CompanyNames)
	})

	t.Run("return one task when one saved", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		res, err := client.GetCrawlingTasks(ctx, &crawler_grpc.GetCrawlingTasksRequest{Site: blindSite})
		require.NoError(t, err)
		require.Equal(t, []string{companyName}, res.CompanyNames)
	})

	t.Run("return multiple tasks when multiple saved", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyNames := []string{"testCompany1", "testCompany2"}
		for _, companyName := range companyNames {
			_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
				CompanyName: companyName,
			})
			require.NoError(t, err)
		}

		res, err := client.GetCrawlingTasks(ctx, &crawler_grpc.GetCrawlingTasksRequest{Site: blindSite})
		require.NoError(t, err)

		require.Equal(t, companyNames, res.CompanyNames)
	})

	t.Run("return empty tasks when all updated", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyNames := []string{"testCompany1", "testCompany2"}
		for _, companyName := range companyNames {
			_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
				CompanyName: companyName,
			})
			require.NoError(t, err)

			_, err = client.SetReviewScore(ctx, &crawler_grpc.SetReviewScoreRequest{
				Site:        blindSite,
				CompanyName: companyName,
				AvgScore:    45,
				ReviewCount: 10,
			})
			require.NoError(t, err)
		}

		res, err := client.GetCrawlingTasks(ctx, &crawler_grpc.GetCrawlingTasksRequest{Site: blindSite})
		require.NoError(t, err)

		require.Empty(t, res.CompanyNames)
	})

	t.Run("return empty tasks when all not exist", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyNames := []string{"testCompany1", "testCompany2"}
		for _, companyName := range companyNames {
			_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
				CompanyName: companyName,
			})
			require.NoError(t, err)

			_, err = client.SetNotExist(ctx, &crawler_grpc.SetNotExistRequest{
				Site:        blindSite,
				CompanyName: companyName,
			})
			require.NoError(t, err)
		}

		res, err := client.GetCrawlingTasks(ctx, &crawler_grpc.GetCrawlingTasksRequest{Site: blindSite})
		require.NoError(t, err)

		require.Empty(t, res.CompanyNames)
	})

	t.Run("return tasks regardless of other review site", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyNames := []string{"testCompany1", "testCompany2"}
		for _, companyName := range companyNames {
			_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
				CompanyName: companyName,
			})
			require.NoError(t, err)

			_, err = client.SetNotExist(ctx, &crawler_grpc.SetNotExistRequest{
				Site:        blindSite,
				CompanyName: companyName,
			})
			require.NoError(t, err)
		}

		res, err := client.GetCrawlingTasks(ctx, &crawler_grpc.GetCrawlingTasksRequest{Site: "otherSite"})
		require.NoError(t, err)

		require.Equal(t, companyNames, res.CompanyNames)
	})

	t.Run("can't update review score if not saved", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)

		_, err := client.SetReviewScore(ctx, &crawler_grpc.SetReviewScoreRequest{
			Site:        blindSite,
			CompanyName: "testCompany",
			AvgScore:    45,
			ReviewCount: 10,
		})
		require.Error(t, err)
	})

	t.Run("can't update review score several times", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		_, err = client.SetReviewScore(ctx, &crawler_grpc.SetReviewScoreRequest{
			Site:        blindSite,
			CompanyName: companyName,
			AvgScore:    45,
			ReviewCount: 10,
		})
		require.NoError(t, err)

		_, err = client.SetReviewScore(ctx, &crawler_grpc.SetReviewScoreRequest{
			Site:        blindSite,
			CompanyName: companyName,
			AvgScore:    45,
			ReviewCount: 10,
		})
		require.Error(t, err)
	})

	t.Run("can't update status not_exist if not saved", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)

		_, err := client.SetNotExist(ctx, &crawler_grpc.SetNotExistRequest{
			Site:        blindSite,
			CompanyName: "testCompany",
		})
		require.Error(t, err)
	})

	t.Run("can't update status not_exist several times", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		_, err = client.SetNotExist(ctx, &crawler_grpc.SetNotExistRequest{
			Site:        blindSite,
			CompanyName: companyName,
		})
		require.NoError(t, err)

		_, err = client.SetNotExist(ctx, &crawler_grpc.SetNotExistRequest{
			Site:        blindSite,
			CompanyName: companyName,
		})
		require.Error(t, err)
	})

	t.Run("can't update different status", func(t *testing.T) {
		t.Run("update not_exist after review score", func(t *testing.T) {
			ctx := context.Background()
			tinit.InitDB(t)
			client := tinit.InitReviewGrpcClient(t)
			providerClient := tinit.InitCrawlingTaskGrpcClient(t)

			companyName := "testCompany"
			_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
				CompanyName: companyName,
			})
			require.NoError(t, err)

			_, err = client.SetNotExist(ctx, &crawler_grpc.SetNotExistRequest{
				Site:        blindSite,
				CompanyName: companyName,
			})
			require.NoError(t, err)

			_, err = client.SetReviewScore(ctx, &crawler_grpc.SetReviewScoreRequest{
				Site:        blindSite,
				CompanyName: companyName,
				AvgScore:    45,
				ReviewCount: 10,
			})
			require.Error(t, err)
		})

		t.Run("update review score after not_exist", func(t *testing.T) {
			ctx := context.Background()
			tinit.InitDB(t)
			client := tinit.InitReviewGrpcClient(t)
			providerClient := tinit.InitCrawlingTaskGrpcClient(t)

			companyName := "testCompany"
			_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
				CompanyName: companyName,
			})
			require.NoError(t, err)

			_, err = client.SetReviewScore(ctx, &crawler_grpc.SetReviewScoreRequest{
				Site:        blindSite,
				CompanyName: companyName,
				AvgScore:    45,
				ReviewCount: 10,
			})
			require.NoError(t, err)

			_, err = client.SetNotExist(ctx, &crawler_grpc.SetNotExistRequest{
				Site:        blindSite,
				CompanyName: companyName,
			})
			require.Error(t, err)
		})
	})

	t.Run("return empty crawling page when nothing saved", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)

		res, err := client.GetCrawlingPages(ctx, &crawler_grpc.GetCrawlingPagesRequest{Site: blindSite})
		require.NoError(t, err)
		require.Empty(t, res.CompanyNames)
	})

	t.Run("return empty crawling page until update score N page", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		res, err := client.GetCrawlingPages(ctx, &crawler_grpc.GetCrawlingPagesRequest{Site: blindSite})
		require.NoError(t, err)
		require.Empty(t, res.CompanyNames)
	})

	t.Run("return crawling page after update score N page", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		companyInfo := &crawler_grpc.SetReviewScoreRequest{
			Site:        blindSite,
			CompanyName: companyName,
			AvgScore:    45,
			ReviewCount: 10,
		}
		_, err = client.SetReviewScore(ctx, companyInfo)
		require.NoError(t, err)

		res, err := client.GetCrawlingPages(ctx, &crawler_grpc.GetCrawlingPagesRequest{Site: blindSite})
		require.NoError(t, err)
		require.Len(t, res.CompanyNames, 1)
		require.Equal(t, companyInfo.CompanyName, res.CompanyNames[0])
	})

	t.Run("return empty crawling page after update not_exist", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		_, err = client.SetNotExist(ctx, &crawler_grpc.SetNotExistRequest{
			Site:        blindSite,
			CompanyName: companyName,
		})
		require.NoError(t, err)

		res, err := client.GetCrawlingPages(ctx, &crawler_grpc.GetCrawlingPagesRequest{Site: blindSite})
		require.NoError(t, err)
		require.Empty(t, res.CompanyNames)
	})

	t.Run("return empty crawling page after update reviewCount zero", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		companyInfo := &crawler_grpc.SetReviewScoreRequest{
			Site:        blindSite,
			CompanyName: companyName,
			AvgScore:    0,
			ReviewCount: 0,
		}
		_, err = client.SetReviewScore(ctx, companyInfo)
		require.NoError(t, err)

		res, err := client.GetCrawlingPages(ctx, &crawler_grpc.GetCrawlingPagesRequest{Site: blindSite})
		require.NoError(t, err)
		require.Empty(t, res.CompanyNames)
	})

	t.Run("return empty crawling page regardless other site", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		companyInfo := &crawler_grpc.SetReviewScoreRequest{
			Site:        blindSite,
			CompanyName: companyName,
			AvgScore:    45,
			ReviewCount: 10,
		}
		_, err = client.SetReviewScore(ctx, companyInfo)
		require.NoError(t, err)

		res, err := client.GetCrawlingPages(ctx, &crawler_grpc.GetCrawlingPagesRequest{Site: "otherSite"})
		require.NoError(t, err)
		require.Empty(t, res.CompanyNames)
	})

	t.Run("save company reviews", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)

		res, err := client.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        blindSite,
			CompanyName: "testCompany",
			Reviews: []*crawler_grpc.Review{
				{
					Score:            45,
					Summary:          "testSummary",
					EmploymentStatus: true,
					ReviewUserId:     "testUserId",
					JobType:          "testJobType",
					UnixMilli:        time.Now().UnixMilli(),
				},
			},
		})
		require.NoError(t, err)
		require.EqualValues(t, 1, res.InsertedCount)
	})

	t.Run("save multiple company reviews", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)

		res, err := client.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        blindSite,
			CompanyName: "testCompany",
			Reviews: []*crawler_grpc.Review{
				{
					Score:            45,
					Summary:          "testSummary",
					EmploymentStatus: true,
					ReviewUserId:     "testUserId",
					JobType:          "testJobType",
					UnixMilli:        time.Now().UnixMilli(),
				},
				{
					Score:            45,
					Summary:          "otherSummary",
					EmploymentStatus: true,
					ReviewUserId:     "otherUserId",
					JobType:          "otherJobType",
					UnixMilli:        time.Now().UnixMilli(),
				},
			},
		})
		require.NoError(t, err)
		require.EqualValues(t, 2, res.InsertedCount)
	})

	t.Run("ignore same company reviews", func(t *testing.T) { //멱등성
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)

		companyName := "testCompany"
		sameReview := &crawler_grpc.Review{
			Score:            45,
			Summary:          "testSummary",
			EmploymentStatus: true,
			ReviewUserId:     "testUserId",
			JobType:          "testJobType",
			UnixMilli:        time.Now().UnixMilli(),
		}

		otherReview := &crawler_grpc.Review{
			Score:            45,
			Summary:          "otherSummary",
			EmploymentStatus: true,
			ReviewUserId:     "otherUserId",
			JobType:          "otherJobType",
			UnixMilli:        time.Now().UnixMilli(),
		}

		res, err := client.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        blindSite,
			CompanyName: companyName,
			Reviews:     []*crawler_grpc.Review{sameReview},
		})
		require.NoError(t, err)
		require.EqualValues(t, 1, res.InsertedCount)

		res, err = client.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        blindSite,
			CompanyName: companyName,
			Reviews:     []*crawler_grpc.Review{sameReview, otherReview},
		})
		require.NoError(t, err)
		require.EqualValues(t, 1, res.InsertedCount)

		res, err = client.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        blindSite,
			CompanyName: companyName,
			Reviews:     []*crawler_grpc.Review{sameReview, otherReview},
		})
		require.NoError(t, err)
		require.EqualValues(t, 0, res.InsertedCount)
	})
}
