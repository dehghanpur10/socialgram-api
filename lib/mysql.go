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
	return user, result.Error
}
func (mySQL *MySQLDatabase) GetUserById(id uint) (*models.User, error) {
	user := new(models.User)
	result := mySQL.DB.First(&user, id)
	return user, result.Error
}
func (mySQL *MySQLDatabase) GetPost(PostId uint) (*models.Post, error) {
	post := new(models.Post)
	result := mySQL.DB.First(&post, PostId)
	return post, result.Error
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
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return errors.New("user don't access to remove")
	}
	return nil
}

func (mySQL *MySQLDatabase) GetLikeStatus(postId, userId uint) (bool, error) {
	var user models.User
	var post models.Post
	post.ID = postId
	err := mySQL.DB.Model(&post).Where("user_id = ?", userId).Association("Likes").Find(&user)
	if err != nil {
		return false, err
	}
	return user.ID == userId, nil
}

func (mySQL *MySQLDatabase) ToggleLike(status bool, postId uint, user *models.User) (bool, error) {
	var err error
	var post models.Post
	post.ID = postId
	if status {
		err = mySQL.DB.Model(&post).Association("Likes").Delete(user)
	} else {
		//err = mySQL.DB.Model(&post).Association("Likes").Append(user)
		query := fmt.Sprintf("INSERT INTO `post_likes` (`post_id`,`user_id`) VALUES (%v,%v) ON DUPLICATE KEY UPDATE `post_id`=`post_id`", postId, user.ID)
		err = mySQL.DB.Exec(query).Error
	}

	if err != nil {
		return status, err
	}
	return !status, nil
}

func (mySQL *MySQLDatabase) IsFriend(user *models.User, friendId uint) (bool, error) {
	var friend models.User
	err := mySQL.DB.Model(user).Where("friend_id = ?", friendId).Association("Friends").Find(&friend)
	if err != nil {
		return false, err
	}
	return friend.ID == friendId, nil
}

func (mySQL *MySQLDatabase) GetFriends(user *models.User) ([]models.User, error) {
	var friends []models.User
	err := mySQL.DB.Model(user).Association("Friends").Find(&friends)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func (mySQL *MySQLDatabase) DeleteFriend(user *models.User, friendId uint) error {
	var friend models.User
	friend.ID = friendId
	err := mySQL.DB.Model(&user).Association("Friends").Delete(&friend)
	if err != nil {
		return err
	}
	return nil
}

func (mySQL *MySQLDatabase) GetFriendsPosts(user *models.User, pageNumber int) ([]models.Post, error) {
	var post []models.Post
	friends, err := mySQL.GetFriends(user)
	if err != nil {
		return nil, err
	}
	result := mySQL.DB.Model(&models.Post{}).Order("id desc").Offset(pageNumber*PAGE_SIZE).Limit(PAGE_SIZE).Where("user_id IN ?", models.ConvertToStringUsers(friends)).Preload("User").Preload("Likes").Find(&post)
	return post, result.Error
}

func (mysql *MySQLDatabase) GetProfileWithUserId(userId uint) (*models.User, error) {
	user := new(models.User)
	result := mysql.DB.Preload("Posts.Likes").First(&user, userId)
	return user, result.Error
}

func (mysql *MySQLDatabase) EditProfile(user *models.User, userInput *models.User) (*models.User, error) {
	result := mysql.DB.Model(&user).Updates(models.User{Bio: userInput.Bio, Interest: userInput.Interest})
	return user, result.Error
}

func (mySQL *MySQLDatabase) GetFollowers(user *models.User) ([]models.User, error) {
	var friendsId []uint
	var friends []models.User
	query := fmt.Sprintf("SELECT user_id FROM user_friends WHERE friend_id = %v", user.ID)
	mySQL.DB.Raw(query).Scan(&friendsId)
	err := mySQL.DB.Find(&friends, friendsId).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}
func (mySQL *MySQLDatabase) CreateRequest(user *models.User, friendId uint) error {
	query := fmt.Sprintf("INSERT INTO `user_requests` (`user_id`,`request_id`) VALUES (%v,%v) ON DUPLICATE KEY UPDATE `user_id`=`user_id`", user.ID, friendId)
	return mySQL.DB.Exec(query).Error
}

func (mySQL *MySQLDatabase) IsRequest(user *models.User, friendId uint) (bool, error) {
	var friend models.User
	err := mySQL.DB.Model(user).Where("request_id = ?", friendId).Association("Requests").Find(&friend)
	if err != nil {
		return false, err
	}
	return friend.ID == friendId, nil
}
