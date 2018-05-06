package graphql

import (
	"family-tree/db"
	t "family-tree/graphql/types"
	"family-tree/utils"

	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func loadParams(params graphql.ResolveParams, key string) interface{} {
	value, isOK := params.Args[key]
	if isOK {
		return value
	}
	return nil
}

func fetchUsersByIDs(IDs []uint64) (resp interface{}, err error) {
	var res []t.User
	var p = bson.M{"id": bson.M{"$in": IDs}}
	err = db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(p).All(&res)

	if err != nil {
		log.Fatal("fetchDetailByUsername: ", err)
	}
	return res, err
}
