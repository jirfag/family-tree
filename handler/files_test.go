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
	r.GET("/files/token?table=user&field=avatar&action=init&table_id=asdas").
		SetDebug(true).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusConflict, r.Code)
		})

}
