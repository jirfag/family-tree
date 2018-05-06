package main

import (
	"family-tree/graphql"
	"family-tree/handler"
	"family-tree/middleware"
	"family-tree/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", middleware.AuthMiddleware.LoginHandler)
	r.POST("/register", handler.RegisterHandler)
	r.POST("/code", handler.VerifyCodeHandler)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		auth.POST("/graphql", graphql.Handler())
		auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
	}

	showStatus()
	r.Run(utils.AppConfig.Server.Host + ":" + utils.AppConfig.Server.Port)

}

func showStatus() {
	fmt.Println("\n===================================" +
		"\nAPP         : " + utils.AppConfig.APPName +
		"\nRunning On  : HTTP      " + utils.AppConfig.Server.Host + ":" + utils.AppConfig.Server.Port +
		"\n===================================")
}
