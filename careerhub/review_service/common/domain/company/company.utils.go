package company

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func RefineNameForSearch(input string) string {
	index := strings.Index(input, "(")
	if index == -1 {
		return input
	}

	return strings.TrimSpace(input[:index])
}

func FilterNotIncludeSite(site string) bson.M {
	return bson.M{ReviewSitesField: bson.M{"$not": bson.M{"$elemMatch": bson.M{SiteField: site}}}}
}

func FilterIncludeSite(site string) bson.M {
	return bson.M{ReviewSitesField: bson.M{"$elemMatch": bson.M{SiteField: site, ExistStatusField: Exist}}}
}
