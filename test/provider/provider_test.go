package provider

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-review-service/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestCrawlingTaskGrpc(t *testing.T) {
	cancelFunc := tinit.RunTestApp(t)
	defer cancelFunc()

	t.Run("return created status when add new crawling task", func(t *testing.T) {
		tinit.InitDB(t)

		ctx := context.Background()
		client := tinit.InitProviderGrpcClient(t)
		res, err := client.AddCrawlingTask(ctx, &provider_grpc.AddCrawlingTaskRequest{
			CompanyName: "testCompany",
		})

		require.NoError(t, err)
		require.Equal(t, "created", res.Status)
	})

	t.Run("return duplicated status when add duplicated crawling task", func(t *testing.T) {
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
		require.Equal(t, "duplicated", res.Status)
	})
}
