package graphql

import (
	t "family-tree/graphql/types"

	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserType",
	Description: "User Type",
	Fields: graphql.Fields{
		"id":               &graphql.Field{Type: graphql.Int},
		"username":         &graphql.Field{Type: graphql.String},
		"realname":         &graphql.Field{Type: graphql.String},
		"email":            &graphql.Field{Type: graphql.String},
		"phone":            &graphql.Field{Type: graphql.String},
		"avatar":           &graphql.Field{Type: graphql.String},
		"wechat":           &graphql.Field{Type: graphql.String},
		"location":         &graphql.Field{Type: graphql.String},
		"verifyCode":       &graphql.Field{Type: graphql.String},
		"createdTime":      &graphql.Field{Type: graphql.String},
		"isGraduated":      &graphql.Field{Type: graphql.Boolean},
		"position":         &graphql.Field{Type: graphql.String},
		"isActivated":      &graphql.Field{Type: graphql.Boolean},
		"isValidated":      &graphql.Field{Type: graphql.Boolean},
		"isBasicCompleted": &graphql.Field{Type: graphql.Boolean},
		"isAdmin":          &graphql.Field{Type: graphql.Boolean},
		"abilities": &graphql.Field{
			Name: "Abilities Type",
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.User).Abilities, nil
			},
		},
		"mentorIDs": &graphql.Field{
			Name: "mentorIDs Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.User).MentorIDs, nil
			},
		},
		"projectIDs": &graphql.Field{
			Name: "projectIDs Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.User).ProjectIDs, nil
			},
		},
		"menteeIDs": &graphql.Field{
			Name: "menteeIDs Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.User).MenteeIDs, nil
			},
		},
		"companyIDs": &graphql.Field{
			Name: "companyIDs Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.User).CompanyIDs, nil
			},
		},
		"groupIDs": &graphql.Field{
			Name: "mentorIDs Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.User).GroupIDs, nil
			},
		},
	},
})

var groupType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GroupType",
	Description: "Group Type",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.Int},
		"groupName":   &graphql.Field{Type: graphql.String},
		"startYear":   &graphql.Field{Type: graphql.String},
		"endYear":     &graphql.Field{Type: graphql.String},
		"createdTime": &graphql.Field{Type: graphql.String},
		"memberIDs": &graphql.Field{
			Name: "memberIDs Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.Group).MemberIDs, nil
			},
		},
	},
})

var projectType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ProjectType",
	Description: "Project Type",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.Int},
		"titile":      &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"year":        &graphql.Field{Type: graphql.Int},
		"startedTime": &graphql.Field{Type: graphql.String},
		"endedTime":   &graphql.Field{Type: graphql.String},
		"github":      &graphql.Field{Type: graphql.String},
		"adminID":     &graphql.Field{Type: graphql.ID},
		"logo":        &graphql.Field{Type: graphql.String},
		"createdTime": &graphql.Field{Type: graphql.String},
		"images": &graphql.Field{
			Name: "images Type",
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.Project).Images, nil
			},
		},
		"files": &graphql.Field{
			Name: "files url Type",
			Type: graphql.NewList(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.Project).Files, nil
			},
		},
		"memberIDs": &graphql.Field{
			Name: "memberIDs Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.Project).MemberIDs, nil
			},
		},
		"roles": &graphql.Field{
			Name: "roles Type",
			Type: graphql.NewList(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(t.Project).Roles, nil
			},
		},
	},
})
