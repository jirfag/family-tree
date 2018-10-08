package types

import (
	"time"
)

// Project is a type to transform project data to database
type Project struct {
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	StartedYear int       `bson:"startedYear" json:"startedYear"`
	EndedYear   int       `bson:"endedYear" json:"endedYear"`
	URL         string    `bson:"url" json:"url"`
	AdminID     uint64    `bson:"adminID" json:"adminID"`
	Github      string    `bson:"github" json:"github"` // github id
	Logo        string    `bson:"logo" json:"logo"`
	Images      []string  `bson:"image" json:"image"` // url of images
	Files       []string  `bson:"files" json:"files"` // url of files
	MemberIDs   []uint64  `bson:"memberIDs" json:"memberIDs"`
	Roles       []string  `bson:"roles" json:"roles"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
}
