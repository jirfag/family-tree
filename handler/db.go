package handler

import (
	"family-tree/db"
	t "family-tree/graphql/types"
	m "family-tree/middleware"
	"family-tree/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitDB(c *gin.Context) {
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").RemoveAll(nil)

	user := t.User{
		ID:          1,
		Username:    utils.AppConfig.Root.Username,
		IsActivated: true,
		IsAdmin:     true,
		InviteCode:  "2333",
		CreatedTime: time.Now(),
	}

	user.Password, _ = m.HashPassword(utils.AppConfig.Root.Password)

	err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Insert(user)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err, "code": http.StatusConflict})
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK", "code": http.StatusOK})
}
