package graphql

import (
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserType",
	Description: "User Type",
	Fields: graphql.Fields{
		"id":               &graphql.Field{Type: graphql.String},
		"password":         &graphql.Field{Type: graphql.String},
		"username":         &graphql.Field{Type: graphql.String},
		"realname":         &graphql.Field{Type: graphql.String},
		"email":            &graphql.Field{Type: graphql.String},
		"phone":            &graphql.Field{Type: graphql.String},
		"avatar":           &graphql.Field{Type: graphql.String},
		"wechat":           &graphql.Field{Type: graphql.String},
		"loaction":         &graphql.Field{Type: graphql.String},
		"inviteCode":       &graphql.Field{Type: graphql.String},
		"createdTime":      &graphql.Field{Type: graphql.String},
		"isGraduate":       &graphql.Field{Type: graphql.Boolean},
		"IsActivated":      &graphql.Field{Type: graphql.Boolean},
		"IsBasicCompleted": &graphql.Field{Type: graphql.Boolean},
		"IsAdmin":          &graphql.Field{Type: graphql.Boolean},
		//"MentorsID"        []int         `bson:"mentorsID" json:"mentorsID"`
		//"MenteesID"        []int         `bson:"menteesID" json:"menteesID"`
		//"GroupsID"         []int         `bson:"groupsID" json:"groupsID"`

		//"detail": &graphql.Field{
		//	Name: "Detail Type",
		//	Type: graphql.NewNonNull(detailType),
		//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//		fmt.Println("Detail Type")
		//		if user, ok := p.Source.(User); ok {
		//			log.Printf("fetching detail of user with username: %s", user.Username)
		//			return fetchDetailByUsername(user.Username)
		//		}
		//		fmt.Println("Detail Type Error")
		//		return nil, nil
		//	},
		//},
	},
})