package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"err":  err,
	})
}
