package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Name      string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Gender    string `gorm:"not null"`
	Age       uint   `gorm:"not null"`
	City      string `gorm:"not null"`
	Country   string `gorm:"not null"`
	AvatarURL string `gorm:"not null"`
	Bio       string
	Interest  string
	Posts      []*Post
	Requests   []*User `gorm:"many2many:User_Request;"`
	Friends   []*User `gorm:"many2many:User_Friend;"`
}
