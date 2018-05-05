package graphql

import (
	"github.com/graphql-go/graphql"
)

// Schema is the GraphQL Schema served by the server.
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
