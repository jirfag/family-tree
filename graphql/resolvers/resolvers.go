package resolvers

import (
	"family-tree/db"
	t "family-tree/graphql/types"
	"family-tree/middleware"
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

func UpdateUser(params graphql.ResolveParams) (interface{}, error) {
	var res t.User
	var p = bson.M{}

	// load params
	username, isOK := params.Args["username"].(string)
	if isOK {
		p["username"] = username
	}
	password, isOK := params.Args["password"].(string)
	if isOK {
		p["password"], _ = middleware.HashPassword(password)
	}
	realname, isOK := params.Args["realname"].(string)
	if isOK {
		p["realname"] = realname
	}
	email, isOK := params.Args["email"].(string)
	if isOK {
		p["email"] = email
	}
	phone, isOK := params.Args["phone"].(string)
	if isOK {
		p["phone"] = phone
	}
	avatar, isOK := params.Args["avatar"].(string)
	if isOK {
		p["avatar"] = avatar
	}
	wechat, isOK := params.Args["wechat"].(string)
	if isOK {
		p["wechat"] = wechat
	}
	loaction, isOK := params.Args["loaction"].(string)
	if isOK {
		p["loaction"] = loaction
	}
	inviteCode, isOK := params.Args["inviteCode"].(string)
	if isOK {
		p["inviteCode"] = inviteCode
	}

	isGraduate, isOK := params.Args["isGraduate"].(bool)
	if isOK {
		p["isGraduate"] = isGraduate
	}
	IsActivated, isOK := params.Args["IsActivated"].(bool)
	if isOK {
		p["IsActivated"] = IsActivated
	}
	IsBasicCompleted, isOK := params.Args["IsBasicCompleted"].(bool)
	if isOK {
		p["IsBasicCompleted"] = IsBasicCompleted
	}
	IsAdmin, isOK := params.Args["IsAdmin"].(bool)
	if isOK {
		p["IsAdmin"] = IsAdmin
	}

	// check user exist
	count, err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": username}).Count()
	if count == 0 || err != nil {
		log.Fatal("UpdateUser: ", err)
	}

	// update user
	err = db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(bson.M{"username": username}, p)
	if err != nil {
		log.Fatal("UpdateUser: ", err)
	}

	// return data
	bsonBytes, _ := bson.Marshal(p)
	bson.Unmarshal(bsonBytes, &res)
	return res, nil
}
