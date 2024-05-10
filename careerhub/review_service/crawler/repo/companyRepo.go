package repo

import (
	"context"

	"github.com/jae2274/careerhub-review-service/common/domain/company"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepo struct {
	col *mongo.Collection
}

func NewCompanyRepo(db *mongo.Database) *CompanyRepo {
	return &CompanyRepo{
		col: db.Collection((&company.Company{}).Collection()),
	}
}

func (r *CompanyRepo) GetCrawlingTasks(ctx context.Context) ([]*company.Company, error) {
	cur, err := r.col.Find(ctx, bson.M{
		company.StatusField: company.Unknown,
	})
	if err != nil {
		return nil, err
	}

	var companies []*company.Company
	if err := cur.All(ctx, &companies); err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *CompanyRepo) SetScoreNPage(ctx context.Context, defaultName string, totalPageCount int32, avgScore int32) (*mongo.UpdateResult, error) {
	filter := bson.M{
		company.DefaultNameField: defaultName,
		company.StatusField:      company.Unknown,
	}
	update := bson.M{
		"$set": bson.M{
			company.CurrentCrawlingPageField: totalPageCount,
			company.AvgScoreField:            avgScore,
			company.StatusField:              company.Exist,
		},
	}
	return r.col.UpdateOne(ctx, filter, update)
}

func (r *CompanyRepo) SetNotExist(ctx context.Context, defaultName string) (*mongo.UpdateResult, error) {
	filter := bson.M{
		company.DefaultNameField: defaultName,
		company.StatusField:      company.Unknown,
	}
	update := bson.M{
		"$set": bson.M{
			company.StatusField: company.NotExist,
		},
	}
	return r.col.UpdateOne(ctx, filter, update)
}
