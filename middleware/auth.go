package middleware

import (
	"family-tree/db"
	t "family-tree/graphql/types"
	"family-tree/utils"
	"log"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// AuthMiddleware is a middleware to validate
var AuthMiddleware = &jwt.GinJWTMiddleware{
	Realm:      "test zone",
	Key:        []byte(utils.AppConfig.Server.SecretKey),
	Timeout:    time.Hour * 24,
	MaxRefresh: time.Hour,
	Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
		var p = bson.M{}
		var res t.User

		err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(p).One(&res)
		if err != nil {
			log.Println("GetUser: ", err)
			return "error", false
		}

		isOK := CheckPasswordHash(password, res.Password)
		if isOK {
			return res.Username, true
		}

		return username, false
	},
	Authorizator: func(userId string, c *gin.Context) bool {
		var p = bson.M{}
		var res t.User

		err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(p).One(&res)
		if err != nil {
			log.Println("GetUser: ", err)
			return false
		}
		return true
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},
	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	TokenLookup: "header:Authorization",
	// TokenLookup: "query:token",
	// TokenLookup: "cookie:token",

	// TokenHeadName is a string in the header. Default value is "Bearer"
	TokenHeadName: "Bearer",

	// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	TimeFunc: time.Now,
}

// CheckPasswordHash is a func to check password hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword is a func to hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
