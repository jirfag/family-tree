package handler

import (
	"family-tree/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck is a func
// @Summary HealthCheck
// @Description Respond pong to ping request
// @Tags additional
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Ping
// @Router /ping [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Ping{Code: http.StatusOK, Message: "pong"})
}
