package graphql

import (
	"github.com/graphql-go/graphql"
)

var queryUser = graphql.Field{
	Name:        "User",
	Description: "Query user",
	Type:        graphql.NewList(userType),

	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{Type: graphql.String},
		"phone":    &graphql.ArgumentConfig{Type: graphql.String},
		"email":    &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: GetUser,
}

// 定义root查询节点
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"user": &queryUser,
	},
})
