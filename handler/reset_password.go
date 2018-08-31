package handler

import (
	"errors"
	"fmt"
	"github.com/fredliang44/family-tree/db"
	"github.com/fredliang44/family-tree/middleware"
	"github.com/fredliang44/family-tree/utils"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"net/http"
)

// GenResetCode is a func to handler register request
// @Summary Gen Reset Code
// @Description Generate register phone sms auth code
// @Tags Reset
// @Accept  json
// @Produce  json
// @Param 	GenResetCode body utils.GenResetCodeReq true "Gen Reset Code"
// @Success 200 {object} utils.StdResp
// @Failure 400 {object} utils.ErrResp
// @Router /reset_password_code [post]
func GenResetCode(c *gin.Context) {
	var info register
	var data register
	c.BindJSON(&info)

	count, err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": info.Username}).Count()

	if count == 0 {
		err = errors.New("No Such User")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrResp{Message: fmt.Sprintln(err), Code: http.StatusBadRequest})
		return
	}

	// update VerifyCode
	rand.Seed(time.Now().Unix())
	code := fmt.Sprintf("%04d", rand.Intn(10000))
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(bson.M{"username": info.Username}, bson.M{"$set": bson.M{"verifyCode": code}})
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": info.Username}).One(&data)

	// send sms
	// Using tencent cloud
	// isOK, msg, errID := utils.SendDYSMS(data.Phone, "SMS_133979618", `{"code":"`+code+`"}`) DAYU example
	isOK, msg, _ := utils.SendQCSMS(data.Phone, 96385, []string{"Family Tree", code})

	if !isOK {
		c.JSON(http.StatusBadRequest, utils.ErrResp{Message: msg, Code: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, utils.StdResp{Message: "OK", Code: http.StatusOK})
}

// ResetPassword is a func to verify sms code
// @Summary Reset Password
// @Description Phrase reset password request
// @Tags Reset
// @Accept  json
// @Produce  json
// @Param 	ResetPassword body utils.ResetReq true "Verify Reset Code"
// @Success 200 {object} utils.VerifyResp
// @Failure 400 {object} utils.ErrResp
// @Router /reset [post]
func ResetPassword(c *gin.Context) {
	var info register
	var data register
	c.BindJSON(&info)

	if info.Username == "" || info.Password == "" || info.VerifyCode == "" {
		c.JSON(http.StatusNotAcceptable, utils.ErrResp{Message: "Missing parameter", Code: http.StatusNotAcceptable})
		return
	}

	// load user info from db
	err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": info.Username}).One(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrResp{Message: "Error occur when fetching user", Code: http.StatusBadRequest})
		return
	}

	if info.VerifyCode != data.VerifyCode {
		c.JSON(http.StatusExpectationFailed, utils.ErrResp{Message: "Wrong code", Code: http.StatusExpectationFailed})
		return
	}

	// Update password
	hashedPassword, _ := middleware.HashPassword(info.Password)
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(
		bson.M{"username": info.Username},
		bson.M{"$set": bson.M{
			"password": hashedPassword}},
	)

	c.JSON(http.StatusOK, utils.StdResp{Message: "OK", Code: http.StatusOK})
}
