package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Password         string        `bson:"password" json:"password"`
	Username         string        `bson:"username" json:"username"`
	RealName         string        `bson:"realName" json:"realName"`
	Email            string        `bson:"email" json:"email"`
	Phone            string        `bson:"phone" json:"phone"`
	Avatar           string        `bson:"avatar" json:"avatar"`
	Wechat           string        `bson:"wechat" json:"wechat"`
	Loaction         string        `bson:"location" json:"location"`
	InviteCode       string        `bson:"inviteCode" json:"inviteCode"`
	CreatedTime      time.Time     `bson:"createdTime" json:"createdTime"`
	IsGraduated      bool          `bson:"isGraduated" json:"isGraduated"`
	IsActivated      bool          `bson:"isActivated" json:"isActivated"`
	IsBasicCompleted bool          `bson:"isBasicCompleted" json:"isBasicCompleted"`
	IsAdmin          bool          `bson:"isAdmin" json:"isAdmin"`
	MentorsID        []int         `bson:"mentorsID" json:"mentorsID"`
	MenteesID        []int         `bson:"menteesID" json:"menteesID"`
	GroupsID         []int         `bson:"groupsID" json:"groupsID"`
}

type Activities struct {
	Action   string `bson:"password" json:"password"`
	Username string `bson:"username" json:"username"`
}
