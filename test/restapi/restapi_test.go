package restapi

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi/restapi_grpc"
	"github.com/jae2274/careerhub-review-service/test/tinit"
	"github.com/jae2274/careerhub-review-service/test/tutils"
	"github.com/stretchr/testify/require"
)

func TestReviewReaderGrpc(t *testing.T) {
	tinit.InitDB(t)
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

		companyScore := &crawler_grpc.SetReviewScoreRequest{
			Site:        site,
			CompanyName: companyName,
			AvgScore:    45,
		}
		_, err = crawlerClient.SetReviewScore(ctx, companyScore)
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyScores(ctx, &restapi_grpc.GetCompanyScoresRequest{
			Site:         site,
			CompanyNames: []string{companyName},
		})

		require.NoError(t, err)
		resultScore, ok := res.CompanyScores[companyName]
		require.True(t, ok)
		require.Equal(t, companyScore.CompanyName, resultScore.CompanyName)
		require.Equal(t, companyScore.AvgScore, resultScore.Score)
		require.Equal(t, companyScore.ReviewCount, resultScore.ReviewCount)
		require.Equal(t, false, resultScore.IsCompleteCrawl)
	})

	t.Run("return iscompletecrawl true when finish crawling task", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()

		site := "testSite"
		companyName := "testCompany"
		_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: companyName,
		})
		require.NoError(t, err)

		companyScore := &crawler_grpc.SetReviewScoreRequest{
			Site:        site,
			CompanyName: companyName,
			AvgScore:    45,
		}
		_, err = crawlerClient.SetReviewScore(ctx, companyScore)
		require.NoError(t, err)
		_, err = crawlerClient.FinishCrawlingTask(ctx, &crawler_grpc.FinishCrawlingTaskRequest{
			Site:        site,
			CompanyName: companyName,
		})
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyScores(ctx, &restapi_grpc.GetCompanyScoresRequest{
			Site:         site,
			CompanyNames: []string{companyName},
		})

		require.NoError(t, err)
		require.Equal(t, true, res.CompanyScores[companyName].IsCompleteCrawl)
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

		companyScore := &crawler_grpc.SetReviewScoreRequest{
			Site:        site,
			CompanyName: companyName, //TODO: 테스트코드 리팩토링 필요
			AvgScore:    45,
		}
		_, err = crawlerClient.SetReviewScore(ctx, companyScore)
		require.NoError(t, err)

		synosymName := "testCompany(주식회사 테스트컴퍼니)"
		res, err := restapiClient.GetCompanyScores(ctx, &restapi_grpc.GetCompanyScoresRequest{
			Site:         site,
			CompanyNames: []string{synosymName},
		})

		require.NoError(t, err)
		resultScore, ok := res.CompanyScores[synosymName]
		require.True(t, ok)
		require.Equal(t, companyScore.CompanyName, resultScore.CompanyName)
		require.Equal(t, companyScore.AvgScore, resultScore.Score)
		require.Equal(t, companyScore.ReviewCount, resultScore.ReviewCount)
		require.Equal(t, false, resultScore.IsCompleteCrawl)
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

		_, err = crawlerClient.SetReviewScore(ctx, &crawler_grpc.SetReviewScoreRequest{
			Site:        site,
			CompanyName: companyName,
			AvgScore:    45,
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

		companyScores := []*crawler_grpc.SetReviewScoreRequest{
			{
				Site:        site,
				CompanyName: "testCompany1",
				AvgScore:    45,
			},
			{
				Site:        site,
				CompanyName: "testCompany2",
				AvgScore:    20,
			},
			{
				Site:        site,
				CompanyName: "testCompany3",
				AvgScore:    35,
			},
		}
		for _, companyScore := range companyScores {
			_, err := providerClient.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
				CompanyName: companyScore.CompanyName,
			})
			require.NoError(t, err)

			_, err = crawlerClient.SetReviewScore(ctx, companyScore)
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
			require.Equal(t, companyScore.CompanyName, resultScore.CompanyName)
			require.Equal(t, companyScore.AvgScore, resultScore.Score)
			require.Equal(t, companyScore.ReviewCount, resultScore.ReviewCount)
			require.Equal(t, false, resultScore.IsCompleteCrawl)
		}
	})

	t.Run("return empty reviews when nothing saved", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()

		res, err := restapiClient.GetCompanyReviews(ctx, &restapi_grpc.GetCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: "testCompany",
		})
		require.NoError(t, err)
		require.Empty(t, res.Reviews)
	})

	t.Run("return empty reviews when not yet saved same company", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()

		_, err := crawlerClient.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: "otherCompany",
			Reviews:     []*crawler_grpc.Review{tutils.NewCompanyReviewReq("", 45, true)},
		})
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyReviews(ctx, &restapi_grpc.GetCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: "testCompany",
		})
		require.NoError(t, err)
		require.Empty(t, res.Reviews)
	})

	t.Run("return reviews when saved", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()

		companyName := "testCompany"
		reviews := []*crawler_grpc.Review{
			tutils.NewCompanyReviewReq("", 45, true),
		}
		_, err := crawlerClient.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: companyName,
			Reviews:     reviews,
		})
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyReviews(ctx, &restapi_grpc.GetCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: companyName,
		})
		require.NoError(t, err)
		tutils.AssertEqualReviews(t, reviews, res.Reviews)
	})

	t.Run("return reviews by synonym company name", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()

		reviews := []*crawler_grpc.Review{
			tutils.NewCompanyReviewReq("", 45, true),
		}
		_, err := crawlerClient.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: "testCompany",
			Reviews:     reviews,
		})
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyReviews(ctx, &restapi_grpc.GetCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: "testCompany(주식회사 테스트컴퍼니)",
		})
		require.NoError(t, err)
		tutils.AssertEqualReviews(t, reviews, res.Reviews)
	})

	t.Run("return multiple reviews after saved", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()

		companyName := "testCompany"
		reviews := []*crawler_grpc.Review{
			tutils.NewCompanyReviewReq("1", 45, true),
			tutils.NewCompanyReviewReq("2", 20, false),
			tutils.NewCompanyReviewReq("3", 35, true),
		}
		_, err := crawlerClient.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: companyName,
			Reviews:     reviews,
		})
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyReviews(ctx, &restapi_grpc.GetCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: companyName,
		})
		require.NoError(t, err)

		slices.Reverse(reviews)
		tutils.AssertEqualReviews(t, reviews, res.Reviews)
	})

	t.Run("returm multiple reviews after saved several times", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()

		companyName := "testCompany"
		reviews1 := []*crawler_grpc.Review{
			tutils.NewCompanyReviewReq("1", 45, true),
			tutils.NewCompanyReviewReq("2", 20, false),
			tutils.NewCompanyReviewReq("3", 35, true),
		}
		reviews2 := []*crawler_grpc.Review{
			tutils.NewCompanyReviewReq("4", 50, true),
			tutils.NewCompanyReviewReq("5", 25, false),
			tutils.NewCompanyReviewReq("6", 30, true),
		}

		reviewsList := [][]*crawler_grpc.Review{reviews1, reviews2}

		for _, reviews := range reviewsList {
			_, err := crawlerClient.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
				Site:        "testSite",
				CompanyName: companyName,
				Reviews:     reviews,
			})
			require.NoError(t, err)
		}

		res, err := restapiClient.GetCompanyReviews(ctx, &restapi_grpc.GetCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: companyName,
		})
		require.NoError(t, err)

		reviews := append(reviews1, reviews2...)
		slices.Reverse(reviews)
		tutils.AssertEqualReviews(t, reviews, res.Reviews)
	})

	t.Run("return reviews sorted by date", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()

		companyName := "testCompany"
		reviews := []*crawler_grpc.Review{
			tutils.NewCompanyReviewReq("1", 45, true),
			tutils.NewCompanyReviewReq("2", 20, false),
			tutils.NewCompanyReviewReq("3", 35, true),
		}
		now := time.Now()
		reviews[1].UnixMilli = now.Add(2 * time.Second).UnixMilli()
		reviews[2].UnixMilli = now.Add(time.Second).UnixMilli()
		reviews[0].UnixMilli = now.UnixMilli()
		_, err := crawlerClient.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: companyName,
			Reviews:     reviews,
		})
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyReviews(ctx, &restapi_grpc.GetCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: companyName,
		})

		require.NoError(t, err)
		expected := []*crawler_grpc.Review{reviews[1], reviews[2], reviews[0]}
		tutils.AssertEqualReviews(t, expected, res.Reviews)
	})

	t.Run("return reviews from offset to limit", func(t *testing.T) {
		tinit.InitDB(t)
		ctx := context.Background()

		companyName := "testCompany"
		reviews := []*crawler_grpc.Review{
			tutils.NewCompanyReviewReq("1", 45, true),
			tutils.NewCompanyReviewReq("2", 20, false),
			tutils.NewCompanyReviewReq("3", 35, true),
			tutils.NewCompanyReviewReq("4", 50, true),
			tutils.NewCompanyReviewReq("5", 25, false),
			tutils.NewCompanyReviewReq("6", 30, true),
		}
		_, err := crawlerClient.SaveCompanyReviews(ctx, &crawler_grpc.SaveCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: companyName,
			Reviews:     reviews,
		})
		require.NoError(t, err)

		res, err := restapiClient.GetCompanyReviews(ctx, &restapi_grpc.GetCompanyReviewsRequest{
			Site:        "testSite",
			CompanyName: companyName,
			Offset:      1,
			Limit:       3,
		})

		require.NoError(t, err)
		expected := []*crawler_grpc.Review{reviews[4], reviews[3], reviews[2]}
		tutils.AssertEqualReviews(t, expected, res.Reviews)
	})
}
