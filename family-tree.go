package main

import (
	"family-tree/graphql"
	"family-tree/handler"
	"family-tree/middleware"
	"family-tree/utils"
	"fmt"
	"github.com/Salvatore-Giordano/gin-redis-ip-limiter"
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	// CORS support
	r.Use(middleware.CORSMiddleware())

	// recovery from internal server error
	r.Use(nice.Recovery(utils.RecoveryHandler))

	// limit request frequency per minute
	r.Use(iplimiter.NewRateLimiterMiddleware(redis.NewClient(&redis.Options{
		Addr:     utils.AppConfig.Redis.Host + ":" + utils.AppConfig.Redis.Port,
		Password: utils.AppConfig.Redis.Password,
		DB:       utils.AppConfig.Redis.DB,
	}), "general", 200, time.Minute))

	// AUTH & Login
	r.POST("/login", middleware.AuthMiddleware.LoginHandler)
	r.POST("/code", handler.GenCode)
	r.POST("/register", handler.RegisterHandler)

	// Healthcheck
	r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong", "code": 200}) })

	// Main Handler
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		auth.POST("/graphql", graphql.Handler())
		auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
		auth.GET("/init_db", handler.InitDB)
	}

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
