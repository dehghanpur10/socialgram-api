package lib

import "gorm.io/gorm"

type MySQLDatabase struct {
	SocialGramStore
	DB *gorm.DB
}