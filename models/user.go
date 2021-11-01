package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string  `gorm:"unique;not null" validate:"required" json:"username"`
	Name      string  `gorm:"not null" validate:"required" json:"name"`
	Email     string  `gorm:"unique;not null" validate:"required" json:"email"`
	Password  string  `gorm:"not null"  json:"password,omitempty"`
	Gender    string  `gorm:"not null" validate:"required" json:"gender"`
	Age       uint    `gorm:"not null" validate:"required" json:"age"`
	City      string  `gorm:"not null" validate:"required" json:"city"`
	Country   string  `gorm:"not null" validate:"required" json:"country"`
	AvatarURL string  `gorm:"not null" json:"image_url,omitempty"`
	Bio       string  `json:"bio"`
	Interest  string  `json:"interest"`
	Posts     []*Post `json:"posts,omitempty"`
	Requests  []*User `gorm:"many2many:User_Request;" json:"requests,omitempty"`
	Friends   []*User `gorm:"many2many:User_Friend;" json:"friends,omitempty"`
}