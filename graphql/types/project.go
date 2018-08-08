package types

import (
	"time"
)

// Project is a type to transform project data to database
type Project struct {
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	StartedYear string    `bson:"startedYear" json:"startedYear "`
	EndedYear   string    `bson:"endedYear" json:"endedYear "`
	AdminID     uint64    `bson:"adminID" json:"adminID"`
	Github      string    `bson:"github" json:"github"`
	Logo        string    `bson:"logo" json:"logo"`
	Images      []string  `bson:"image" json:"image"`
	Files       []string  `bson:"files" json:"files"` // url
	MemberIDs   []uint64  `bson:"memberIDs" json:"memberIDs"`
	Roles       []string  `bson:"roles" json:"roles"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
}
