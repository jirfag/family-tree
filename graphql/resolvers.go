package graphql

import (
	"errors"
	"family-tree/db"
	t "family-tree/graphql/types"
	"family-tree/middleware"
	"family-tree/utils"
	"log"

	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// GetUser is a graphql resolver to get user info
func GetUser(params graphql.ResolveParams) (interface{}, error) {
	var res []t.User
	var p = bson.M{}

	id, isOK := params.Args["id"].(uint64)
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
		log.Println("GetUser: ", err)
		return nil, nil
	}
	return res, nil
}

// GetUser is a graphql resolver to get user info
func GetGroup(params graphql.ResolveParams) (interface{}, error) {
	var res []t.Group
	var p = bson.M{}

	id, isOK := params.Args["id"].(uint64)
	if isOK {
		p["id"] = id
	}

	groupName, isOK := params.Args["groupName"].(string)
	if isOK {
		p["groupNamesername"] = groupName
	}
	startYear, isOK := params.Args["startYear"].(int)
	if isOK {
		p["startYear"] = startYear
	}
	endYear, isOK := params.Args["endYear"].(int)
	if isOK {
		p["endYear"] = endYear
	}

	err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(p).All(&res)

	if err != nil {
		log.Println("Get Group: ", err)
		return nil, nil
	}
	return res, nil
}

// UpdateUser is a graphql resolver to update user info
func UpdateUser(params graphql.ResolveParams) (interface{}, error) {

	var res t.User
	var p = bson.M{}

	// load params
	username, isOK := params.Args["username"].(string)
	if isOK && username != "" {
		p["username"] = username
	}

	// load params
	id, isOK := params.Args["username"].(uint64)
	if isOK && id != 0 {
		p["id"] = id
	}

	// check user exist
	err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(p).One(&p)
	if err != nil {
		log.Println("Update User: ", err)
	}

	if params.Context.Value("User") == username {
		// load data
		id, isOK := params.Args["id"].(bson.ObjectId)
		if isOK {
			p["id"] = id
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

		// update user
		err = db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(bson.M{"username": username}, p)
		if err != nil {
			log.Fatal("Update User: ", err)
		}

		// return data
		bsonBytes, _ := bson.Marshal(p)
		bson.Unmarshal(bsonBytes, &res)
		return res, nil
	}
	return nil, errors.New("You Couldn't Change other's info")
}
