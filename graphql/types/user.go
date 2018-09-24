package types

import (
	"time"
)

// User is a type to transform user data to database
type User struct {
	ID               uint64    `json:"id" bson:"_id,omitempty"`
	Password         string    `bson:"password" json:"password"`
	Username         string    `bson:"username" json:"username"`
	RealName         string    `bson:"realName" json:"realName"`
	Email            string    `bson:"email" json:"email"`
	Phone            string    `bson:"phone" json:"phone"`
	Avatar           string    `bson:"avatar" json:"avatar"`
	Gender           bool      `bson:"gender" json:"gender"`
	Wechat           string    `bson:"wechat" json:"wechat"`     // Wechat ID
	Location         string    `bson:"location" json:"location"` // Current Location (City)
	VerifyCode       string    `bson:"verifyCode" json:"verifyCode"`
	CreatedTime      time.Time `bson:"createdTime" json:"createdTime"`           // Created time
	JoinedYear       int       `bson:"joinedYear" json:"joinedYear"`             // When joined Unique Studio
	EnrollmentYear   int       `bson:"enrollmentYear" json:"enrollmentYear"`     // When entered school
	Position         string    `bson:"position" json:"position"`                 // Position in company
	IsGraduated      bool      `bson:"isGraduated" json:"isGraduated"`           // Whether graduated
	IsActivated      bool      `bson:"isActivated" json:"isActivated"`           // Account was activated via verify code
	IsValidated      bool      `bson:"isValidated" json:"isValidated"`           // Account was validated by admin
	IsBasicCompleted bool      `bson:"isBasicCompleted" json:"isBasicCompleted"` // Whether account basic info completed
	IsAdmin          bool      `bson:"isAdmin" json:"isAdmin"`
	Abilities        []string  `bson:"abilities" json:"abilities"` // User Abilities; EX: Python
	ProjectIDs       []uint64  `bson:"projectIDs" json:"projectIDs"`
	MentorIDs        []uint64  `bson:"mentorIDs" json:"mentorIDs"`
	MenteeIDs        []uint64  `bson:"menteeIDs" json:"menteeIDs"`
	CompanyIDs       []uint64  `bson:"CompanyIDs" json:"CompanyIDs"`
	GroupIDs         []uint64  `bson:"groupIDs" json:"groupIDs"`
}
