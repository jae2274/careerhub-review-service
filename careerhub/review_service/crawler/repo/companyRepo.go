package repo

import (
	"context"
	"fmt"

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

func filterNotIncludeSite(site string) bson.M {
	return bson.M{company.ReviewSitesField: bson.M{"$not": bson.M{"$elemMatch": bson.M{company.SiteField: site}}}}
}

func (r *CompanyRepo) GetCrawlingTasks(ctx context.Context, site string) ([]*company.Company, error) {
	cur, err := r.col.Find(ctx, filterNotIncludeSite(site))
	if err != nil {
		return nil, err
	}

	var companies []*company.Company
	if err := cur.All(ctx, &companies); err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *CompanyRepo) SetScoreNPage(ctx context.Context, defaultName string, reviewSite *company.ReviewSite) error {
	filter := filterNotIncludeSite(reviewSite.Site)
	filter[company.DefaultNameField] = defaultName

	update := bson.M{
		"$push": bson.M{company.ReviewSitesField: reviewSite},
	}

	result, err := r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("matched count is 0. company: %s, site: %s", defaultName, reviewSite.Site)
	}

	return nil
}

func (r *CompanyRepo) SetNotExist(ctx context.Context, defaultName string, site string) error {
	filter := filterNotIncludeSite(site)
	filter[company.DefaultNameField] = defaultName

	reviewSite := &company.ReviewSite{
		Site:   site,
		Status: company.NotExist,
	}
	update := bson.M{
		"$push": bson.M{company.ReviewSitesField: reviewSite},
	}

	result, err := r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("matched count is 0. company: %s, site: %s", defaultName, site)
	}

	return nil
}
