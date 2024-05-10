package company

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DefaultNameField = "defaultName"
	OtherNamesField  = "otherNames"
	ReviewSitesField = "reviewSites"
	ReviewSites_Site = "reviewSites.site"
	SiteField        = "site"
	StatusField      = "status"
	InsertedAtField  = "insertedAt"
	UpdatedAtField   = "updatedAt"
)

type Company struct {
	DefaultName string        `bson:"defaultName"`
	OtherNames  []string      `bson:"otherNames"`
	ReviewSites []*ReviewSite `bson:"reviewSites"`
	InsertedAt  time.Time     `bson:"insertedAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
}

type ReviewSite struct {
	Site                string `bson:"site"`
	Status              Status `bson:"status"`
	AvgScore            int32  `bson:"avgScore"`
	CurrentCrawlingPage int32  `bson:"currentCrawlingPage"`
}

func (*Company) Collection() string {
	return "company"
}

func (*Company) IndexModels() map[string]*mongo.IndexModel {
	indexName := fmt.Sprintf("%s_1", DefaultNameField)
	return map[string]*mongo.IndexModel{
		indexName: {
			Keys:    bson.D{{Key: DefaultNameField, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
}
