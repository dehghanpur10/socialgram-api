package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID   uint    `gorm:"unique;not null"`
	Title    string  `gorm:"unique;not null"`
	Content  string  `gorm:"unique;not null"`
	ImageURL string  `gorm:"not null"`
	Likes    []*User `gorm:"many2many:Post_Like;"`
}
