package db

import (
	"family-tree/utils"

	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
)

var DBSession = initDB()

func initDB() *mgo.Session {

	info := &mgo.DialInfo{
		Addrs:    []string{utils.AppConfig.Mongo.Host + ":" + utils.AppConfig.Mongo.Port},
		Database: utils.AppConfig.Mongo.DB,
		Username: utils.AppConfig.Mongo.Username,
		Password: utils.AppConfig.Mongo.Password,
	}

	session, err := mgo.DialWithInfo(info)
	ai.Connect(session.DB(utils.AppConfig.Mongo.DB).C("user"))
	ai.Connect(session.DB(utils.AppConfig.Mongo.DB).C("project"))
	ai.Connect(session.DB(utils.AppConfig.Mongo.DB).C("group"))

	if err != nil {
		panic(err)
	}
	//defer session.Close()

	return session
}
