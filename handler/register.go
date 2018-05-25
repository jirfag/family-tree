package handler

import (
	"errors"
	"family-tree/db"
	"family-tree/middleware"
	"family-tree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
	"math/rand"
	"net/http"
	"time"
)

// GenCode is a func to gen register code
// @Summary Gen Register Code
// @Description Generate register phone sms auth code
// @Tags Register
// @Accept  json
// @Produce  json
// @Param 	GenCode body utils.GenCodeReq true "Gen Register Code"
// @Success 200 {object} utils.StdResp
// @Failure 400 {object} utils.ErrResp
// @Router /register_code [post]
func GenCode(c *gin.Context) {
	var info register
	c.BindJSON(&info)

	count, err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": info.Username}).Count()
	if count != 0 {
		err = errors.New("Username exists")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrResp{Message: fmt.Sprintln(err), Code: http.StatusConflict})
		return
	}

	// log CreatedTime
	info.CreatedTime = time.Now()

	// gen verifyCode
	rand.Seed(time.Now().Unix())
	info.VerifyCode = fmt.Sprintf("%04d", rand.Intn(10000))

	// handle password
	if err := validator.Validate(info); err != nil {
		c.JSON(http.StatusNotAcceptable, utils.ErrResp{Message: fmt.Sprintln(err), Code: http.StatusNotAcceptable})
		return
	}

	info.Password, _ = middleware.HashPassword(info.Password)

	// send sms
	// Using tencent cloud
	// isOK, msg, errID := utils.SendDYSMS(data.Phone, "SMS_133979618", `{"code":"`+code+`"}`) DAYU example
	isOK, msg, _ := utils.SendQCSMS(info.Phone, 96385, []string{"Family Tree", info.VerifyCode})

	if !isOK {
		c.JSON(http.StatusBadRequest, utils.ErrResp{Message: msg, Code: http.StatusBadRequest})
		return
	}

	// save user info
	info.ID = ai.Next("user")
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Insert(info)

	c.JSON(http.StatusOK, utils.StdResp{Message: "OK", Code: http.StatusOK})
}

// RegisterHandler is a func to verify sms code
// @Summary Register
// @Description Register after verify sms code
// @Tags Register
// @Accept  json
// @Produce  json
// @Param 	Register body utils.RegisterReq true "Verify Register Code"
// @Success 200 {object} utils.VerifyResp
// @Failure 400 {object} utils.ErrResp
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	var data register
	var info register
	c.BindJSON(&data)

	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": data.Username}).One(&info)

	if info.Username == "" {
		c.JSON(http.StatusNotFound, utils.ErrResp{Message: "No such user", Code: http.StatusNotFound})
		return
	}

	if info.IsActivated == true && info.VerifyCode == data.VerifyCode {
		c.JSON(http.StatusConflict, utils.ErrResp{Message: "Already Activated", Code: http.StatusConflict})
		return
	}

	if info.VerifyCode == data.VerifyCode {
		info.IsActivated = true
		db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(bson.M{"username": data.Username}, &info)
		c.JSON(http.StatusOK, utils.StdResp{Message: "OK", Code: http.StatusOK})
		return
	}
	c.JSON(http.StatusForbidden, utils.ErrResp{Message: "Wrong Verify Code", Code: http.StatusForbidden})
}

type register struct {
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	Username    string    `validate:"min=3,max=20" bson:"username" json:"username"`
	Password    string    `validate:"min=3,max=30,regexp=^((?=.*\d)|(?=.*\W+))(?![.\n])(?=.*[a-z]).*$" bson:"password" json:"password"`
	Phone       string    `validate:"len=11,regexp=^1[34578]\d{9}$" bson:"phone" json:"phone"`
	VerifyCode  string    `bson:"verifyCode" json:"verifyCode"`
	IsActivated bool      `bson:"isActivated" json:"isActivated"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
}
