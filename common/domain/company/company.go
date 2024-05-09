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
)

type Company struct {
	DefaultName string    `bson:"defaultName"`
	OtherNames  []string  `bson:"otherNames"`
	InsertedAt  time.Time `bson:"insertedAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
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
