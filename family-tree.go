package main

import (
	"family-tree/graphql"
	"family-tree/handler"
	"family-tree/middleware"
	"family-tree/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS support
	r.Use(middleware.CORSMiddleware())

	// AUTH & Login
	r.POST("/login", middleware.AuthMiddleware.LoginHandler)
	r.POST("/register", handler.RegisterHandler)
	r.POST("/code", handler.VerifyCodeHandler)

	// Main Handler
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		auth.POST("/graphql", graphql.Handler())
		auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
	}

	// Healthcheck
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "code": 200})
	})

	// Show Status
	showStatus()

	// Run Server
	r.Run(utils.AppConfig.Server.Host + ":" + utils.AppConfig.Server.Port)
}

func showStatus() {
	fmt.Println("\n===================================" +
		"\nAPP         : " + utils.AppConfig.APPName +
		"\nRunning On  : HTTP      " + utils.AppConfig.Server.Host + ":" + utils.AppConfig.Server.Port +
		"\n===================================")
}
