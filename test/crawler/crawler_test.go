package crawler

import (
	"context"
	"testing"

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

			_, err = client.SetScoreNPage(ctx, &crawler_grpc.SetScoreNPageRequest{
				Site:           blindSite,
				CompanyName:    companyName,
				AvgScore:       45,
				TotalPageCount: 10,
				PageSize:       15,
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

	t.Run("can't update score and pageCount if not saved", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)

		_, err := client.SetScoreNPage(ctx, &crawler_grpc.SetScoreNPageRequest{
			Site:           blindSite,
			CompanyName:    "testCompany",
			AvgScore:       45,
			TotalPageCount: 10,
			PageSize:       15,
		})
		require.Error(t, err)
	})

	t.Run("can't update score and pageCount several times", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		_, err = client.SetScoreNPage(ctx, &crawler_grpc.SetScoreNPageRequest{
			Site:           blindSite,
			CompanyName:    companyName,
			AvgScore:       45,
			TotalPageCount: 10,
			PageSize:       15,
		})
		require.NoError(t, err)

		_, err = client.SetScoreNPage(ctx, &crawler_grpc.SetScoreNPageRequest{
			Site:           blindSite,
			CompanyName:    companyName,
			AvgScore:       45,
			TotalPageCount: 10,
			PageSize:       15,
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
		t.Run("update not_exist after score and pageCount", func(t *testing.T) {
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

			_, err = client.SetScoreNPage(ctx, &crawler_grpc.SetScoreNPageRequest{
				Site:           blindSite,
				CompanyName:    companyName,
				AvgScore:       45,
				TotalPageCount: 10,
				PageSize:       15,
			})
			require.Error(t, err)
		})

		t.Run("update score and pageCount after not_exist", func(t *testing.T) {
			ctx := context.Background()
			tinit.InitDB(t)
			client := tinit.InitReviewGrpcClient(t)
			providerClient := tinit.InitCrawlingTaskGrpcClient(t)

			companyName := "testCompany"
			_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
				CompanyName: companyName,
			})
			require.NoError(t, err)

			_, err = client.SetScoreNPage(ctx, &crawler_grpc.SetScoreNPageRequest{
				Site:           blindSite,
				CompanyName:    companyName,
				AvgScore:       45,
				TotalPageCount: 10,
				PageSize:       15,
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
		require.Empty(t, res.CrawlingPages)
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
		require.Empty(t, res.CrawlingPages)
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

		companyInfo := &crawler_grpc.SetScoreNPageRequest{
			Site:           blindSite,
			CompanyName:    companyName,
			AvgScore:       45,
			TotalPageCount: 10,
			PageSize:       15,
		}
		_, err = client.SetScoreNPage(ctx, companyInfo)
		require.NoError(t, err)

		res, err := client.GetCrawlingPages(ctx, &crawler_grpc.GetCrawlingPagesRequest{Site: blindSite})
		require.NoError(t, err)
		require.Len(t, res.CrawlingPages, 1)
		require.Equal(t, companyInfo.TotalPageCount, res.CrawlingPages[0].CurrentCrawlingPage)
		require.Equal(t, companyInfo.PageSize, res.CrawlingPages[0].PageSize)
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
		require.Empty(t, res.CrawlingPages)
	})

	t.Run("return empty crawling page after update pageCount zero", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitReviewGrpcClient(t)
		providerClient := tinit.InitCrawlingTaskGrpcClient(t)

		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		companyInfo := &crawler_grpc.SetScoreNPageRequest{
			Site:           blindSite,
			CompanyName:    companyName,
			AvgScore:       45,
			TotalPageCount: 0,
			PageSize:       15,
		}
		_, err = client.SetScoreNPage(ctx, companyInfo)
		require.NoError(t, err)

		res, err := client.GetCrawlingPages(ctx, &crawler_grpc.GetCrawlingPagesRequest{Site: blindSite})
		require.NoError(t, err)
		require.Empty(t, res.CrawlingPages)
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

		companyInfo := &crawler_grpc.SetScoreNPageRequest{
			Site:           blindSite,
			CompanyName:    companyName,
			AvgScore:       45,
			TotalPageCount: 10,
			PageSize:       15,
		}
		_, err = client.SetScoreNPage(ctx, companyInfo)
		require.NoError(t, err)

		res, err := client.GetCrawlingPages(ctx, &crawler_grpc.GetCrawlingPagesRequest{Site: "otherSite"})
		require.NoError(t, err)
		require.Empty(t, res.CrawlingPages)
	})
}
