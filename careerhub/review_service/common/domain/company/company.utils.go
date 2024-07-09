package company

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func RefineNameForSearch(input string) string {
	fitstBracket := strings.Index(input, "(")
	if fitstBracket == -1 {
		return input
	}
	secondBracket := strings.Index(input, ")")
	if secondBracket == -1 {
		return input
	}

	return strings.TrimSpace(input[:fitstBracket]) + strings.TrimSpace(input[secondBracket+1:])
}

func FilterNotIncludeSite(site string) bson.M {
	return bson.M{ReviewSitesField: bson.M{"$not": bson.M{"$elemMatch": bson.M{SiteField: site}}}}
}

func FilterIncludeSite(site string) bson.M {
	return bson.M{ReviewSitesField: bson.M{"$elemMatch": bson.M{SiteField: site, ExistStatusField: Exist}}}
}
