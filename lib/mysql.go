package lib

import (
	"gorm.io/gorm"
	"socialgram/models"
)

type MySQLDatabase struct {
	SocialGramStore
	DB *gorm.DB
}

func (mySQL *MySQLDatabase) CreateNewUser(user *models.User) error {
	return mySQL.DB.Model(&models.User{}).Create(user).Error
}

func (mySQL *MySQLDatabase) GetUser(username string) (*models.User, error) {
	user := new(models.User)
	result := mySQL.DB.Where("username = ?", username).First(&user)
	return user, result.Error
}
