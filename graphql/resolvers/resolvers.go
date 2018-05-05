package resolvers

import (
	"family-tree/db"
	t "family-tree/graphql/types"
	"family-tree/utils"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func GetUser(params graphql.ResolveParams) (interface{}, error) {

	var res []t.User
	var p = bson.M{}
	id, isOK := params.Args["id"].(string)
	if isOK {
		p["id"] = id
	}
	username, isOK := params.Args["username"].(string)
	if isOK {
		p["username"] = username
	}
	phone, isOK := params.Args["phone"].(string)
	if isOK {
		p["phone"] = phone
	}
	school, isOK := params.Args["email"].(string)
	if isOK {
		p["email"] = school
	}

	err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(p).All(&res)

	if err != nil {
		log.Fatal("GetUser: ", err)
	}

	return res, nil
}
