package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID   uint    `gorm:"not null"  json:"user_id"`
	User	 *User    `gorm:"-"  json:"user"`
	Title    string  `gorm:"not null" validate:"required" json:"title"`
	Content  string  `gorm:"not null" validate:"required" json:"content"`
	ImageURL string  `gorm:"not null" json:"image_url"`
	Likes    []*User `gorm:"many2many:Post_Like;constraint:OnUpdate:CASCADE" json:"likes,omitempty"`
}
