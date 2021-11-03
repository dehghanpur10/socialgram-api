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
}
