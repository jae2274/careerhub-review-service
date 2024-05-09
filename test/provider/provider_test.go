package provider

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/server"
	"github.com/jae2274/careerhub-review-service/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestCrawlingTaskGrpc(t *testing.T) {
	cancelFunc := tinit.RunTestApp(t)
	defer cancelFunc()

	t.Run("return created when add new crawling task", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()
		client := tinit.InitProviderGrpcClient(t)
		res, err := client.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: "testCompany",
		})

		require.NoError(t, err)
		require.Equal(t, server.CrawlingTaskCreated, res.Status)
	})

	t.Run("return duplicated when add duplicated crawling task", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()
		client := tinit.InitProviderGrpcClient(t)

		_, err := client.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: "testCompany",
		})
		require.NoError(t, err)

		// Add duplicated crawling task
		res, err := client.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: "testCompany",
		})

		require.NoError(t, err)
		require.Equal(t, server.CrawlingTaskDuplicated, res.Status)
	})

	//이음동의어의 자세한 조건은 별도의 테스트로 분리
	t.Run("return not_modified when has synonym company", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()
		client := tinit.InitProviderGrpcClient(t)

		_, err := client.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: "testCompany",
		})
		require.NoError(t, err)

		// Add synonym company
		res, err := client.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: "testCompany(주)",
		})

		require.NoError(t, err)
		require.Equal(t, server.CrawlingTaskNotModified, res.Status)
	})
}
