package repo

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/domain/company"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CompanyRepo struct {
	col *mongo.Collection
}

func NewCompanyRepo(db *mongo.Database) *CompanyRepo {
	return &CompanyRepo{
		col: db.Collection((&company.Company{}).Collection()),
	}
}

func (cr *CompanyRepo) FindCompany(ctx context.Context, defaultName string, originName string) (*company.Company, bool, error) {
	filter := bson.M{
		company.DefaultNameField: defaultName,
		company.OtherNamesField:  originName,
	}

	var c company.Company
	err := cr.col.FindOne(ctx, filter).Decode(&c)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &c, true, nil
}

func (cr *CompanyRepo) Save(ctx context.Context, defaultName string, originName string) (*mongo.UpdateResult, error) {
	// Upsert 필터: defaultName을 기준으로 문서를 찾는다.
	filter := bson.M{company.DefaultNameField: defaultName}

	// 업데이트 내용: $setOnInsert는 문서가 삽입될 때만 적용되며, $addToSet은 항상 적용된다.
	update := bson.M{
		"$setOnInsert": bson.M{
			company.DefaultNameField: defaultName,
			company.InsertedAtField:  time.Now(),
			company.ReviewSitesField: []*company.ReviewSite{},
		},
		"$addToSet": bson.M{company.OtherNamesField: originName},
		"$set":      bson.M{company.UpdatedAtField: time.Now()},
	}

	// Upsert 옵션을 활성화한다.
	opts := options.Update().SetUpsert(true)

	// 컬렉션에서 upsert 실행
	return cr.col.UpdateOne(ctx, filter, update, opts)

}
