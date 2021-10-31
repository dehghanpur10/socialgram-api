package lib

import "socialgram/models"

type SocialGramStore interface {
	 CreateNewUser(user *models.User) error
}
