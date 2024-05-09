package company

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const NameField = "name"

type Company struct {
	Name       string    `bson:"name"`
	InsertedAt time.Time `bson:"insertedAt"`
}

func (*Company) Collection() string {
	return "company"
}

func (*Company) IndexModels() map[string]*mongo.IndexModel {
	indexName := fmt.Sprintf("%s_1", NameField)
	return map[string]*mongo.IndexModel{
		indexName: {
			Keys:    bson.D{{Key: NameField, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
}
