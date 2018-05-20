package handler

import (
	"family-tree/db"
	"family-tree/middleware"
	"family-tree/utils"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"net/http"
)

// GenResetCode is a func to handler register request
func GenResetCode(c *gin.Context) {
	var info register
	c.BindJSON(&info)

	// update InviteCode
	rand.Seed(time.Now().Unix())
	code := fmt.Sprintf("%04d", rand.Intn(10000))
	err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(bson.M{"username": info.Username}, bson.M{"$set": bson.M{"inviteCode": code}})

	// send sms
	if os.Getenv("GIN_MODE") == "release" {
		isOK, msg, errID := utils.SendSMS(info.Phone, "SMS_133979618", `{"code":"`+code+`"}`)
		if !isOK {
			c.JSON(http.StatusBadRequest, gin.H{"message": msg, "err_id": errID})
			return
		}
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// ResetPassword is a func to verify sms code
func ResetPassword(c *gin.Context) {
	var info register
	var data register
	c.BindJSON(&info)

	if info.Username == "" || info.Password == "" || info.InviteCode == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Missing parameter", "code": http.StatusNotAcceptable})
		return
	}

	// load user info from db
	err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": info.Username}).One(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error occur when fetching user", "code": http.StatusBadRequest})
	}
	if info.InviteCode != data.InviteCode {
		c.JSON(http.StatusExpectationFailed, gin.H{"message": "Wrong code", "code": http.StatusExpectationFailed})
		return
	}

	// Update password
	hashedPassword, _ := middleware.HashPassword(info.Password)
	err = db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(
		bson.M{"username": info.Username},
		bson.M{"$set": bson.M{
			"password": hashedPassword}},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error occur when fetching user", "code": http.StatusBadRequest})
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "code": http.StatusOK})
}
