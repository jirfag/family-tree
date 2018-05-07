package graphql

import (
	"context"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

// Handler is a func to handler graphql requests
func Handler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	return func(c *gin.Context) {
		// Load user
		userID := jwt.ExtractClaims(c)["id"]

		ctx := c.Request.Context()

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")

		h.ServeHTTP(c.Writer, c.Request.WithContext(context.WithValue(ctx, "User", userID)))
	}
}
