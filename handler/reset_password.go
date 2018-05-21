package handler

import (
	"errors"
	"family-tree/db"
	"family-tree/middleware"
	"family-tree/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"net/http"
)

// GenResetCode is a func to handler register request
func GenResetCode(c *gin.Context) {
	var info register
	var data register
	c.BindJSON(&info)

	count, err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": info.Username}).Count()

	if count == 0 {
		err = errors.New("No Such User")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintln(err)})
		return
	}

	// update VerifyCode
	rand.Seed(time.Now().Unix())
	code := fmt.Sprintf("%04d", rand.Intn(10000))
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(bson.M{"username": info.Username}, bson.M{"$set": bson.M{"verifyCode": code}})
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": info.Username}).One(&data)

	// send sms
	isOK, msg, errID := utils.SendSMS(data.Phone, "SMS_133979618", `{"code":"`+code+`"}`)
	if !isOK {
		c.JSON(http.StatusBadRequest, gin.H{"message": msg, "err_id": errID})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "code": http.StatusOK})
}

// ResetPassword is a func to verify sms code
func ResetPassword(c *gin.Context) {
	var info register
	var data register
	c.BindJSON(&info)

	if info.Username == "" || info.Password == "" || info.VerifyCode == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Missing parameter", "code": http.StatusNotAcceptable})
		return
	}

	// load user info from db
	err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": info.Username}).One(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error occur when fetching user", "code": http.StatusBadRequest})
		return
	}

	if info.VerifyCode != data.VerifyCode {
		c.JSON(http.StatusExpectationFailed, gin.H{"message": "Wrong code", "code": http.StatusExpectationFailed})
		return
	}

	// Update password
	hashedPassword, _ := middleware.HashPassword(info.Password)
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(
		bson.M{"username": info.Username},
		bson.M{"$set": bson.M{
			"password": hashedPassword}},
	)

	c.JSON(http.StatusOK, gin.H{"message": "OK", "code": http.StatusOK})
}
