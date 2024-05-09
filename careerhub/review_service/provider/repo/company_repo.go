package repo

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-review-service/common/domain/company"
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

func (c *CompanyRepo) AddCompany(ctx context.Context, company *company.Company) error {
	company.InsertedAt = time.Now()
	_, err := c.col.InsertOne(ctx, company)

	return err
}
