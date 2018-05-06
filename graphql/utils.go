package graphql

import (
	"github.com/graphql-go/graphql"
)

func loadParams(params graphql.ResolveParams, key string) interface{} {
	value, isOK := params.Args[key]
	if isOK {
		return value
	}
	return nil
}
