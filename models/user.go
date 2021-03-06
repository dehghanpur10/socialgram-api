package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string  `gorm:"index;unique;not null" validate:"required" json:"username"`
	Name            string  `gorm:"not null" validate:"required" json:"name"`
	Email           string  `gorm:"unique;not null" validate:"required,email" json:"email"`
	Password        string  `gorm:"not null"  json:"password,omitempty"`
	Gender          string  `gorm:"not null" validate:"required" json:"gender"`
	Age             uint    `gorm:"not null" validate:"required" json:"age,omitempty"`
	City            string  `gorm:"not null" validate:"required" json:"city"`
	Country         string  `gorm:"not null" validate:"required" json:"country"`
	AvatarURL       string  `gorm:"not null" validate:"required" json:"image_url"`
	Bio             string  `json:"bio"`
	Interest        string  `json:"interest"`
	Request         bool    `json:"is_requested"`
	IsFriend        bool    `json:"isFriend"`
	FollowerNumber  int     `json:"follower_number"`
	FollowingNumber int     `json:"following_number"`
	Posts           []*Post `json:"posts,omitempty"`
	Requests        []*User `gorm:"many2many:User_Request;constraint:OnUpdate:CASCADE;" json:"requests,omitempty"`
	Friends         []*User `gorm:"many2many:User_Friend;constraint:OnUpdate:CASCADE;" json:"friends,omitempty"`
}
