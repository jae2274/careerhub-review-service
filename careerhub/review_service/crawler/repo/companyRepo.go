package repo

import (
	"context"
	"fmt"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/domain/company"
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

func (r *CompanyRepo) GetCrawlingTasks(ctx context.Context, site string) ([]*company.Company, error) {
	cur, err := r.col.Find(ctx, company.FilterNotIncludeSite(site))
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
	filter := company.FilterNotIncludeSite(reviewSite.Site)
	filter[company.DefaultNameField] = defaultName

	reviewSite.ExistStatus = company.Exist
	reviewSite.CrawlingStatus = company.NotCrawled
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
	filter := company.FilterNotIncludeSite(site)
	filter[company.DefaultNameField] = defaultName

	reviewSite := &company.ReviewSite{
		Site:           site,
		ExistStatus:    company.NotExist,
		CrawlingStatus: company.NotCrawled,
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

func (r *CompanyRepo) GetCrawlingTargets(ctx context.Context, site string) ([]*company.Company, error) {
	filter := bson.M{company.ReviewSitesField: bson.M{"$elemMatch": bson.M{company.SiteField: site, company.ExistStatusField: company.Exist, company.CrawlingStatusField: company.NotCrawled, company.ReviewCountField: bson.M{"$gt": 0}}}}
	cur, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var companies []*company.Company
	if err := cur.All(ctx, &companies); err != nil {
		return nil, err
	}

	return companies, nil
}
