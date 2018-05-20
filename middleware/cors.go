package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var CORSConfig = loadConfig()

func loadConfig() cors.Config {
	Origins := []string{}
	if os.Getenv("GIN_MODE") == "release" {
		Origins = []string{"*"}
	} else {
		Origins = []string{"*"}
	}
	config := cors.Config{
		AllowOrigins:     Origins,
		AllowMethods:     []string{"GET", "PUT", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			fmt.Println(origin)
			return true
			//return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}
	return config
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO uncomment when release
		//if os.Getenv("GIN_MODE") == "release" {
		//	c.Writer.Header().Set("Access-Control-Allow-Origin", "https://fmt.fredliang.cn")
		//} else {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
