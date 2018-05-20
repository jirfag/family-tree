package handler

import (
	"family-tree/utils"
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGenCode(t *testing.T) {
	r := gofight.New()
	r.POST("/register_code").
		SetDebug(true).
		SetJSON(gofight.D{
			"username": "test_family_tree",
			"password": "test",
			"phone":    "17777766666",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, "{\"message\":\"OK\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestRegister(t *testing.T) {
	r := gofight.New()

	resp_list := []string{"{\"code\":200,\"message\":\"OK\",\"status\":\"Already Verifyed\"}",
		"{\"code\":200,\"message\":\"OK\",\"status\":\"Verifyed\"}"}
	r.POST("/register").
		SetDebug(true).
		SetJSON(gofight.D{
			"username": utils.AppConfig.Root.Username,
			"password": "2333",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Contains(t, resp_list, r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
