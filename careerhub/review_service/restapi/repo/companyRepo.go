package repo

import (
	"context"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/domain/company"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepo struct {
	col *mongo.Collection
}

func NewCompanyRepo(col *mongo.Database) *CompanyRepo {
	return &CompanyRepo{
		col: col.Collection((&company.Company{}).Collection()),
	}
}

func (r *CompanyRepo) GetCompanies(ctx context.Context, site string, companyNames []string) ([]*company.Company, error) {
	filter := company.FilterIncludeSite(site)
	filter[company.DefaultNameField] = bson.M{"$in": companyNames}

	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var companies []*company.Company
	if err := cursor.All(ctx, &companies); err != nil {
		return nil, err
	}
	return companies, nil
}
