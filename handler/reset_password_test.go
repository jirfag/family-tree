package handler

import (
	"family-tree/utils"
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGenResetCode(t *testing.T) {
	r := gofight.New()

	r.POST("/reset_password_code").
		SetDebug(true).
		SetJSON(gofight.D{
			"username": "test_family_tree",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, "{\"code\":200, \"message\":\"OK\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestResetPassword(t *testing.T) {
	r := gofight.New()

	r.POST("/reset").
		SetDebug(true).
		SetJSON(gofight.D{
			"username":   utils.AppConfig.Root.Username,
			"password":   utils.AppConfig.Root.Password,
			"verifyCode": "2333",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, "{\"code\":200, \"message\":\"OK\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
