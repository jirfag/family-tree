package handler

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestGetPolicyToken(t *testing.T) {
	r := gofight.New()

	// check err sending code
	r.GET("/files/token?table=user&field=avatar&action=init&table_id=0").
		SetDebug(true).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})

	// check table error
	r.GET("/files/token?table=user&field=avatar&action=init&table_id=asdas").
		SetDebug(true).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusConflict, r.Code)
		})

	// check get id error
	r.GET("/files/token?table=user&field=avatar&action=init&table_id=asdas").
		SetDebug(true).
		SetHeader(gofight.H{
			"Authorization": "Bear ",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusConflict, r.Code)
		})

}

func TestFilesCallBack(t *testing.T) {
	r := gofight.New()

	// check err sending code
	r.GET("/files/callback").
		SetDebug(true).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotFound, r.Code)
		})

	// check err sending code
	r.POST("/files/callback").
		SetDebug(true).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotAcceptable, r.Code)
		})

}
