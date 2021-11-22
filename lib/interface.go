package lib

import "socialgram/models"

type SocialGramStore interface {
	CreateNewUser(user *models.User) error
	GetUser(userInfo string) (*models.User, error)
	GetPost(PostId uint) (*models.Post, error)
	SearchUsers(userInfo string, pageNumber int) ([]models.User, error)
	CreateNewPost(post *models.Post) error
	DeletePost(postId,userId uint) error
	GetLikeStatus(postId, userId uint) (bool,error)
	ToggleLike(status bool,postId uint,user *models.User) (bool, error)
	IsFriend(userId *models.User,friendId uint) (bool,error)
	GetFriends(user *models.User) ([]models.User,error)
	DeleteFriend(user *models.User, friendId  uint) error
	GetFriendsPosts(user *models.User, pageNumber int) ([]models.Post, error)
	GetProfileWithUserId(userId uint) (*models.User, error)
	EditProfile(user *models.User, userInput *models.User) (*models.User, error)
	GetFollowers(user *models.User) ([]models.User,error)
}
