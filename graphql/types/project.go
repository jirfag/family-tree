package types

import (
	"time"
)

type Projects struct {
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	Year        int       `bson:"year" json:"year"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
	AdminID     string    `bson:"adminID" json:"adminID"`
	Logo        string    `bson:"logo" json:"logo"`
	Image       []string  `bson:"image" json:"image"`
	members     []User    `bson:"members" json:"members"`
}
