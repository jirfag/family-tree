package handler

import (
	"family-tree/db"
	"family-tree/middleware"
	"family-tree/utils"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func RegisterHandler(c *gin.Context) {
	var info register
	c.BindJSON(&info)

	count, err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": info.Username}).Count()
	if count != 0 {
		c.JSON(http.StatusConflict, gin.H{"msg": "Username exists."})
		return
	} else {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}
	}
	rand.Seed(time.Now().Unix())
	info.InviteCode = fmt.Sprintf("%04d", rand.Intn(10000))
	info.Password, err = middleware.HashPassword(info.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	utils.SendSMS(info.Phone, "SMS_133979618", `{"code":"`+info.InviteCode+`"}`)
	err = db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Insert(info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func VerifyCodeHandler(c *gin.Context) {
	var data register
	var info register
	c.BindJSON(&data)

	err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": data.Username}).One(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if info.IsActivated == true {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "status": "Already Verifyed"})
		return
	}

	if info.InviteCode == data.InviteCode {
		info.IsActivated = true
		db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(bson.M{"username": data.Username}, &info)
		c.JSON(http.StatusOK, gin.H{"message": "OK", "status": "Verifyed"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "Wrong Verify Code"})
		return
	}

}

type register struct {
	Username    string `bson:"username" json:"username"`
	Password    string `bson:"password" json:"password"`
	Phone       string `bson:"phone" json:"phone"`
	InviteCode  string `bson:"inviteCode" json:"inviteCode"`
	IsActivated bool   `bson:"isActivated" json:"isActivated"`
}
