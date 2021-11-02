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

func (mySQL *MySQLDatabase) GetUser(userInfo string) (*models.User, error) {
	user := new(models.User)
	result := mySQL.DB.Where("username = ? OR email = ?", userInfo, userInfo).First(&user)
	return user, result.Error
}

func (mySQL *MySQLDatabase) SearchUsers(userInfo string, pageNumber int) ([]models.User, error) {
	var users []models.User
	userInfo = "%" + userInfo + "%"
	result := mySQL.DB.Offset(pageNumber * PAGE_SIZE).Limit(PAGE_SIZE).Where("username LIKE ? OR name LIKE ?", userInfo, userInfo).Find(&users)
	return users, result.Error
}
