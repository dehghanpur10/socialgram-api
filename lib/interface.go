package lib

import "socialgram/models"

type SocialGramStore interface {
	CreateNewUser(user *models.User) error
	GetUser(userInfo string) (*models.User, error)
	SearchUsers(userInfo string, pageNumber int) ([]models.User, error)
	CreateNewPost(post *models.Post) error
	DeletePost(postId,userId uint) error

}
