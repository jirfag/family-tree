package resolvers

import (
	"family-tree/db"
	t "family-tree/graphql/types"
	"family-tree/utils"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

func loadParams(params graphql.ResolveParams, key string) interface{} {
	value, isOK := params.Args[key]
	if isOK {
		return value
	}
	return nil
}
