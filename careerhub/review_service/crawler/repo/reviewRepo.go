package repo

import (
	"context"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/domain/review"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReviewRepo struct {
	col *mongo.Collection
}

func NewReviewRepo(db *mongo.Database) *ReviewRepo {
	return &ReviewRepo{
		col: db.Collection((&review.Review{}).Collection()),
	}
}

func (r *ReviewRepo) InsertReviews(ctx context.Context, rv []*review.Review) (int, error) {
	docs := make([]interface{}, 0, len(rv))
	for _, v := range rv {
		docs = append(docs, v)
	}

	result, err := r.col.InsertMany(ctx, docs, options.InsertMany().SetOrdered(false)) //멱등성을 위해 ordered를 false로 설정

	if err != nil && !mongo.IsDuplicateKeyError(err) {
		return len(result.InsertedIDs), err
	}

	return len(result.InsertedIDs), nil
}
