package middleware

import (
	"github.com/gin-contrib/cors"
	"os"
	"time"
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
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
			//return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}
	return config
}
