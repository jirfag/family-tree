package types

import (
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User Type",
	Description: "User Type of FMT",
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

type User struct {
	ID               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Password         string        `bson:"password" json:"password"`
	Username         string        `bson:"username" json:"username"`
	RealName         string        `bson:"realName" json:"realName"`
	Email            string        `bson:"email" json:"email"`
	Phone            string        `bson:"phone" json:"phone"`
	Avatar           string        `bson:"avatar" json:"avatar"`
	Wechat           string        `bson:"wechat" json:"wechat"`
	Loaction         string        `bson:"location" json:"location"`
	InviteCode       string        `bson:"inviteCode" json:"inviteCode"`
	CreatedTime      string        `bson:"createdTime" json:"createdTime"`
	IsGraduated      string        `bson:"isGraduated" json:"isGraduated"`
	IsActivated      string        `bson:"isActivated" json:"isActivated"`
	IsBasicCompleted string        `bson:"isBasicCompleted" json:"isBasicCompleted"`
	IsAdmin          string        `bson:"isAdmin" json:"isAdmin"`
	MentorsID        []int         `bson:"mentorsID" json:"mentorsID"`
	MenteesID        []int         `bson:"menteesID" json:"menteesID"`
	GroupsID         []int         `bson:"groupsID" json:"groupsID"`
}

type Activities struct {
	Action   string `bson:"password" json:"password"`
	Username string `bson:"username" json:"username"`
}
