package review

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SiteField             = "site"
	CompanyNameField      = "companyName"
	ScoreField            = "score"
	SummaryField          = "summary"
	EmploymentStatusField = "employmentStatus"
	ReviewUserIdField     = "reviewUserId"
	JobTypeField          = "jobType"
	DateField             = "date"
)

type Review struct {
	Site             string    `bson:"site"`
	CompanyName      string    `bson:"companyName"`
	Score            int32     `bson:"score"`
	Summary          string    `bson:"summary"`
	EmploymentStatus bool      `bson:"employmentStatus"`
	ReviewUserId     string    `bson:"reviewUserId"`
	JobType          string    `bson:"jobType"`
	Date             time.Time `bson:"date"`
}

func (*Review) Collection() string {
	return "review"
}

func (*Review) IndexModels() map[string]*mongo.IndexModel {
	indexName := fmt.Sprintf("%s_1_%s_1_%s_1_%s_1", SiteField, CompanyNameField, SummaryField, ReviewUserIdField)

	return map[string]*mongo.IndexModel{
		indexName: {
			Keys:    bson.D{{Key: SiteField, Value: 1}, {Key: CompanyNameField, Value: 1}, {Key: SummaryField, Value: 1}, {Key: ReviewUserIdField, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
}
