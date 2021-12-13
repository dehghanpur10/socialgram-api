package appHandler

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"socialgram/lib"
	"socialgram/models"
)

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - GetProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - GetProfileHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}

	userIdProfile, err := lib.GetUserIdFromQuery(r)
	if err != nil {
		fmt.Println("GetUserIdFromQuery - GetProfileHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}
	if userIdProfile == 0 {
		userIdProfile = user.ID
	}

	resultUser, err := db.GetProfileWithUserId(userIdProfile)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(" GetProfileWithUserId - GetProfileHandler error:", err)
			lib.HttpError404(w, "user not found ")
			return
		}
		fmt.Println("GetFriendsPosts - GetProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}

	resultUser.Password = ""
	isRequest, err := db.IsRequest(user, userIdProfile)
	if err != nil {
		fmt.Println("IsRequest - GetProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}
	if isRequest {
		resultUser.Request = true
	} else {
		resultUser.Request = false
	}
	isFriend, err := db.IsFriend(user, userIdProfile)
	if err != nil {
		fmt.Println("IsFriend - GetProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}
	if !isFriend && (userIdProfile != user.ID) {
		resultUser.Posts = []*models.Post{}
	}
	if isFriend {
		resultUser.IsFriend = true
	}
	resultUsers, err := db.GetFollowers(user)
	if err != nil {
		fmt.Println("EditProfile - GetProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}

	friends, err := db.GetFriends(user)
	if err != nil {
		fmt.Println("GetFriends - GetProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}

	resultUser.FollowingNumber = len(friends)
	resultUser.FollowerNumber = len(resultUsers)

	jsonBytes, err := lib.ConvertToJsonBytes(resultUser)
	if err != nil {
		fmt.Println("json.Marshal - GetProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)

}
