package restapi

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi/restapi_grpc"
	"github.com/jae2274/careerhub-review-service/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestReviewReaderGrpc(t *testing.T) {
	cancelFunc := tinit.RunTestApp(t)
	defer cancelFunc()

	restapiClient := tinit.InitReviewReaderGrpcClient(t)
	providerClient := tinit.InitCrawlingTaskGrpcClient(t)
	crawlerClient := tinit.InitReviewGrpcClient(t)

	t.Run("return empty companyScore when nothing saved", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()

		site := "testSite"
		companyName := "testCompany"
		res, err := restapiClient.GetCompanyScores(ctx, &restapi_grpc.GetCompanyScoresRequest{
			Site:         site,
			CompanyNames: []string{companyName},
		})

		require.NoError(t, err)
		require.Empty(t, res.CompanyScores)
	})

	t.Run("return empty companyScore when not yet updated", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()

		site := "testSite"
		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyScores(ctx, &restapi_grpc.GetCompanyScoresRequest{
			Site:         site,
			CompanyNames: []string{companyName},
		})

		require.NoError(t, err)
		require.Empty(t, res.CompanyScores)
	})

	t.Run("return companyScore when updated", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()

		site := "testSite"
		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		companyScore := &crawler_grpc.SetScoreNPageRequest{
			Site:           site,
			CompanyName:    companyName,
			AvgScore:       45,
			TotalPageCount: 10,
		}
		_, err = crawlerClient.SetScoreNPage(ctx, companyScore)
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyScores(ctx, &restapi_grpc.GetCompanyScoresRequest{
			Site:         site,
			CompanyNames: []string{companyName},
		})

		require.NoError(t, err)
		resultScore, ok := res.CompanyScores[companyName]
		require.True(t, ok)
		require.Equal(t, companyScore.AvgScore, resultScore)
	})
	t.Run("return companyScore by synonym name", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()

		site := "testSite"
		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		companyScore := &crawler_grpc.SetScoreNPageRequest{
			Site:           site,
			CompanyName:    companyName, //TODO: 테스트코드 리팩토링 필요
			AvgScore:       45,
			TotalPageCount: 10,
		}
		_, err = crawlerClient.SetScoreNPage(ctx, companyScore)
		require.NoError(t, err)

		synosymName := "testCompany(주식회사 테스트컴퍼니)"
		res, err := restapiClient.GetCompanyScores(ctx, &restapi_grpc.GetCompanyScoresRequest{
			Site:         site,
			CompanyNames: []string{synosymName},
		})

		require.NoError(t, err)
		resultScore, ok := res.CompanyScores[synosymName]
		require.True(t, ok)
		require.Equal(t, companyScore.AvgScore, resultScore)
	})

	t.Run("return empty companyScore when updated status not_exist", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()

		site := "testSite"
		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		_, err = crawlerClient.SetNotExist(ctx, &crawler_grpc.SetNotExistRequest{
			Site:        site,
			CompanyName: companyName,
		})
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyScores(ctx, &restapi_grpc.GetCompanyScoresRequest{
			Site:         site,
			CompanyNames: []string{companyName},
		})

		require.NoError(t, err)
		require.Empty(t, res.CompanyScores)
	})

	t.Run("return empty companyScore by other site", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()

		site := "testSite"
		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		_, err = crawlerClient.SetScoreNPage(ctx, &crawler_grpc.SetScoreNPageRequest{
			Site:           site,
			CompanyName:    companyName,
			AvgScore:       45,
			TotalPageCount: 10,
		})
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyScores(ctx, &restapi_grpc.GetCompanyScoresRequest{
			Site:         "otherSite",
			CompanyNames: []string{companyName},
		})

		require.NoError(t, err)
		require.Empty(t, res.CompanyScores)
	})

	t.Run("return multiple companyScores", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()

		site := "testSite"

		companyScores := []*crawler_grpc.SetScoreNPageRequest{
			{
				Site:           site,
				CompanyName:    "testCompany1",
				AvgScore:       45,
				TotalPageCount: 14,
			},
			{
				Site:           site,
				CompanyName:    "testCompany2",
				AvgScore:       20,
				TotalPageCount: 10,
			},
			{
				Site:           site,
				CompanyName:    "testCompany3",
				AvgScore:       35,
				TotalPageCount: 14,
			},
		}
		for _, companyScore := range companyScores {
			_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
				CompanyName: companyScore.CompanyName,
			})
			require.NoError(t, err)

			_, err = crawlerClient.SetScoreNPage(ctx, companyScore)
			require.NoError(t, err)
		}

		companyNames := make([]string, 0, len(companyScores))
		for _, companyScore := range companyScores {
			companyNames = append(companyNames, companyScore.CompanyName)
		}

		res, err := restapiClient.GetCompanyScores(ctx, &restapi_grpc.GetCompanyScoresRequest{
			Site:         site,
			CompanyNames: companyNames,
		})

		require.NoError(t, err)
		for _, companyScore := range companyScores {
			resultScore, ok := res.CompanyScores[companyScore.CompanyName]
			require.True(t, ok)
			require.Equal(t, companyScore.AvgScore, resultScore)
		}
	})
}
