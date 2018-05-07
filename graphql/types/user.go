package types

import (
	"time"
)

type User struct {
	ID               uint64    `json:"id" bson:"_id,omitempty"`
	Password         string    `bson:"password" json:"password"`
	Username         string    `bson:"username" json:"username"`
	RealName         string    `bson:"realName" json:"realName"`
	Email            string    `bson:"email" json:"email"`
	Phone            string    `bson:"phone" json:"phone"`
	Avatar           string    `bson:"avatar" json:"avatar"`
	Wechat           string    `bson:"wechat" json:"wechat"`
	Loaction         string    `bson:"location" json:"location"`
	InviteCode       string    `bson:"inviteCode" json:"inviteCode"`
	CreatedTime      time.Time `bson:"createdTime" json:"createdTime"`
	IsGraduated      bool      `bson:"isGraduated" json:"isGraduated"`
	IsActivated      bool      `bson:"isActivated" json:"isActivated"`
	IsBasicCompleted bool      `bson:"isBasicCompleted" json:"isBasicCompleted"`
	IsAdmin          bool      `bson:"isAdmin" json:"isAdmin"`
	Mentors          []User    `bson:"mentors" json:"mentors"`
	Mentees          []User    `bson:"mentees" json:"mentees"`
	Groups           []Group   `bson:"groups" json:"groups"`
	MentorIDs        []uint64  `bson:"mentorIDs" json:"mentorIDs"`
	MenteeIDs        []uint64  `bson:"menteeIDs" json:"menteeIDs"`
	GroupIDs         []uint64  `bson:"groupIDs" json:"groupIDs"`
}
