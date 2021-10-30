package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null" validate:"required" json:"username"`
	Name      string `gorm:"unique;not null" validate:"required" json:"name"`
	Email     string `gorm:"unique;not null" validate:"required" json:"email"`
	Password  string `gorm:"not null" validate:"required" json:"password,omitempty"`
	Gender    string `gorm:"not null" validate:"required" json:"gender"`
	Age       uint   `gorm:"not null" validate:"required" json:"age"`
	City      string `gorm:"not null" validate:"required" json:"city"`
	Country   string `gorm:"not null" validate:"required" json:"country"`
	AvatarURL string `gorm:"not null" validate:"required" json:"avatar_url"`
	Bio       string `json:"bio,omitempty"`
	Interest  string `json:"interest,omitempty"`
	Posts     []*Post
	Requests  []*User `gorm:"many2many:User_Request;"`
	Friends   []*User `gorm:"many2many:User_Friend;"`
}
