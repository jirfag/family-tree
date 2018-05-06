package graphql

import (
	"family-tree/graphql/resolvers"
	t "family-tree/graphql/types"
	"github.com/graphql-go/graphql"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"updateUser": &mutateUser,
		//"status": &queryStatus,
	},
})

var mutateUser = graphql.Field{
	Name:        "User",
	Description: "Mutate user",
	Type:        graphql.NewList(t.UserType), // the return type for this field
	// Args是定义在GraphQL查询中支持的查询字段，
	// 可自行随意定义，如加上limit,start这类
	Args: graphql.FieldConfigArgument{
		"id":               &graphql.ArgumentConfig{Type: graphql.String},
		"password":         &graphql.ArgumentConfig{Type: graphql.String},
		"username":         &graphql.ArgumentConfig{Type: graphql.String},
		"realname":         &graphql.ArgumentConfig{Type: graphql.String},
		"email":            &graphql.ArgumentConfig{Type: graphql.String},
		"phone":            &graphql.ArgumentConfig{Type: graphql.String},
		"avatar":           &graphql.ArgumentConfig{Type: graphql.String},
		"wechat":           &graphql.ArgumentConfig{Type: graphql.String},
		"loaction":         &graphql.ArgumentConfig{Type: graphql.String},
		"inviteCode":       &graphql.ArgumentConfig{Type: graphql.String},
		"createdTime":      &graphql.ArgumentConfig{Type: graphql.String},
		"isGraduate":       &graphql.ArgumentConfig{Type: graphql.Boolean},
		"IsActivated":      &graphql.ArgumentConfig{Type: graphql.Boolean},
		"IsBasicCompleted": &graphql.ArgumentConfig{Type: graphql.Boolean},
		"IsAdmin":          &graphql.ArgumentConfig{Type: graphql.Boolean},
	},
	// Resolve是一个处理请求的函数，具体处理逻辑可在此进行
	Resolve: resolvers.UpdateUser,
}
