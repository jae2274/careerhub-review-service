package company

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DefaultNameField               = "defaultName"
	OtherNamesField                = "otherNames"
	ReviewSitesField               = "reviewSites"
	ReviewSites_Site               = "reviewSites.site"
	SiteField                      = "site"
	ExistStatusField               = "existStatus"
	ReviewSite_CrawlingStatusField = "reviewSites.crawlingStatus"
	CrawlingStatusField            = "crawlingStatus"
	ReviewCountField               = "reviewCount"
	InsertedAtField                = "insertedAt"
	UpdatedAtField                 = "updatedAt"
)

type Company struct {
	DefaultName string        `bson:"defaultName"`
	OtherNames  []string      `bson:"otherNames"`
	ReviewSites []*ReviewSite `bson:"reviewSites"`
	InsertedAt  time.Time     `bson:"insertedAt"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
}

type ReviewSite struct {
	Site           string         `bson:"site"`
	ExistStatus    ExistStatus    `bson:"existStatus"`
	AvgScore       int32          `bson:"avgScore"`
	ReviewCount    int32          `bson:"reviewCount"`
	TotalPageCount int32          `bson:"totalPageCount"`
	CrawlingStatus CrawlingStatus `bson:"crawlingStatus"`
}

func (*Company) Collection() string {
	return "company"
}

func (*Company) IndexModels() map[string]*mongo.IndexModel {
	defaultNameIndex := fmt.Sprintf("%s_1", DefaultNameField)
	otherNamesIndex := fmt.Sprintf("%s_1", OtherNamesField)
	return map[string]*mongo.IndexModel{
		defaultNameIndex: {
			Keys:    bson.D{{Key: DefaultNameField, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		otherNamesIndex: {
			Keys:    bson.D{{Key: OtherNamesField, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
}
