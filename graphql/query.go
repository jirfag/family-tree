package graphql

import (
	"family-tree/graphql/resolvers"
	t "family-tree/graphql/types"
	"github.com/graphql-go/graphql"
)

var queryUser = graphql.Field{
	Name:        "User",
	Description: "Query user",
	Type:        graphql.NewList(t.UserType),
	// Args是定义在GraphQL查询中支持的查询字段，
	// 可自行随意定义，如加上limit,start这类
	Args: graphql.FieldConfigArgument{
		"id":       &graphql.ArgumentConfig{Type: graphql.String},
		"username": &graphql.ArgumentConfig{Type: graphql.String},
		"phone":    &graphql.ArgumentConfig{Type: graphql.String},
		"email":    &graphql.ArgumentConfig{Type: graphql.String},
	},
	// Resolve是一个处理请求的函数，具体处理逻辑可在此进行
	Resolve: resolvers.GetUser,
}

// 定义root查询节点
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"user": &queryUser,
		//"status": &queryStatus,
	},
})
