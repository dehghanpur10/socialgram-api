package lib

import (
	"errors"
	"fmt"
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
	fmt.Println(result.RowsAffected)
	return user, result.Error
}

func (mySQL *MySQLDatabase) SearchUsers(userInfo string, pageNumber int) ([]models.User, error) {
	var users []models.User
	userInfo = "%" + userInfo + "%"
	result := mySQL.DB.Offset(pageNumber*PAGE_SIZE).Limit(PAGE_SIZE).Where("username LIKE ? OR name LIKE ?", userInfo, userInfo).Find(&users)
	return users, result.Error
}
func (mySQL *MySQLDatabase) CreateNewPost(post *models.Post) error {
	return mySQL.DB.Model(&models.Post{}).Create(post).Error
}

func (mySQL *MySQLDatabase) DeletePost(postId, userId uint) error {
	result := mySQL.DB.Where("user_id = ? AND deleted_at IS NULL", userId).Delete(&models.Post{}, postId)
	if result.Error != nil{
		return  result.Error
	}else if  result.RowsAffected < 1 {
		return errors.New("this post not found or user don't access to remove")
	}
	return nil
}
