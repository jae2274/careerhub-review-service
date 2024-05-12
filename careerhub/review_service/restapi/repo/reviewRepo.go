package repo

import (
	"context"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/domain/review"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *ReviewRepo) GetReviews(ctx context.Context, site string, companyName string, offset int64, limit int64) ([]*review.Review, error) {
	filter := bson.M{
		review.SiteField:        site,
		review.CompanyNameField: companyName,
	}

	opt := options.Find().SetSort(bson.M{review.DateField: -1}).SetSkip(offset).SetLimit(limit)
	cursor, err := r.col.Find(ctx, filter, opt)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []*review.Review{}, nil
		}
		return nil, err
	}

	var reviews []*review.Review
	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}
