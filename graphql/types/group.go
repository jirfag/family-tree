package types

import (
	"time"
)

type Group struct {
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	GroupName   string    `bson:"groupName" json:"groupName"`
	StartYear   int       `bson:"startYear" json:"startYear"`
	EndYear     int       `bson:"endYear" json:"endYear"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
	Members     []User    `bson:"members" json:"members"`
	MemberIDs   []uint64  `bson:"memberIDs" json:"memberIDs"`
}
