package tutils

import (
	"testing"
	"time"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi/restapi_grpc"
	"github.com/stretchr/testify/require"
)

func NewCompanyReviewReq(prefix string, score int32, status bool) *crawler_grpc.Review {
	time.Sleep(100 * time.Millisecond)
	return &crawler_grpc.Review{
		Score:            score,
		Summary:          prefix + "Summary",
		EmploymentStatus: status,
		ReviewUserId:     prefix + "ReviewUserId",
		JobType:          prefix + "JobType",
		UnixMilli:        time.Now().UnixMilli(),
	}
}

func AssertEqualReviews(t *testing.T, expected []*crawler_grpc.Review, actual []*restapi_grpc.Review) {
	require.Len(t, actual, len(expected))
	for i, e := range expected {
		AssertEqualReview(t, e, actual[i])
	}
}

func AssertEqualReview(t *testing.T, expected *crawler_grpc.Review, actual *restapi_grpc.Review) {
	require.Equal(t, expected.Score, actual.Score)
	require.Equal(t, expected.Summary, actual.Summary)
	require.Equal(t, expected.EmploymentStatus, actual.EmploymentStatus)
	require.Equal(t, expected.ReviewUserId, actual.ReviewUserId)
	require.Equal(t, expected.JobType, actual.JobType)
	require.Equal(t, expected.UnixMilli, actual.UnixMilli)
}
