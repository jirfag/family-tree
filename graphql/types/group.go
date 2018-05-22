package types

import (
	"time"
)

// Group is a type to transform group data to database
type Group struct {
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	GroupName   string    `bson:"groupName" json:"groupName"`
	StartYear   int       `bson:"startYear" json:"startYear"`
	EndYear     int       `bson:"endYear" json:"endYear"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
	MemberIDs   []uint64  `bson:"memberIDs" json:"memberIDs"`
}
