package handler

import (
	"family-tree/db"
	"family-tree/middleware"
	"family-tree/utils"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2/bson"
	validator "gopkg.in/validator.v2"
)

// RegisterHandler is a func to handler register request
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

	// log CreatedTime
	info.CreatedTime = time.Now()

	// gen InviteCode
	rand.Seed(time.Now().Unix())
	info.InviteCode = fmt.Sprintf("%04d", rand.Intn(10000))

	// handle password
	if err := validator.Validate(info); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err})
		return
	}
	info.Password, err = middleware.HashPassword(info.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	// send sms
	if os.Getenv("GIN_MODE") == "release" {
		isOK, msg, errID := utils.SendSMS(info.Phone, "SMS_133979618", `{"code":"`+info.InviteCode+`"}`)
		if !isOK {
			c.JSON(http.StatusBadRequest, gin.H{"message": msg, "err_id": errID})
			return
		}
	}

	// save user info
	info.ID = ai.Next("user")
	err = db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Insert(info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// VerifyCodeHandler is a func to verify sms code
func VerifyCodeHandler(c *gin.Context) {
	var data register
	var info register
	c.BindJSON(&data)

	err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": data.Username}).One(&info)
	fmt.Println("info.Username:", info.Username)

	if info.Username == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No such user"})
		return
	}

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
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	Username    string    `validate:"min=3,max=20" bson:"username" json:"username"`
	Password    string    `validate:"min=3,max=30,regexp=^((?=.*\d)|(?=.*\W+))(?![.\n])(?=.*[a-z]).*$" bson:"password" json:"password"`
	Phone       string    `validate:"len=11,regexp=^1[34578]\d{9}$" bson:"phone" json:"phone"`
	InviteCode  string    `bson:"inviteCode" json:"inviteCode"`
	IsActivated bool      `bson:"isActivated" json:"isActivated"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
}
