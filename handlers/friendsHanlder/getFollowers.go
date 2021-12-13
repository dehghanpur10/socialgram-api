package friendsHanlder

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"socialgram/lib"
	"socialgram/models"
)

func GetFollowersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - GetFollowersHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - GetFollowersHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}

	userIdProfile, err := lib.GetUserIdFromQuery(r)
	if err != nil {
		fmt.Println("GetUserIdFromQuery - GetFollowersHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}
	if userIdProfile == 0 {
		userIdProfile = user.ID
	}
	isFriend, err := db.IsFriend(user, userIdProfile)
	if err != nil {
		fmt.Println("IsFriend - GetFollowersHandler error:", err)
		lib.HttpError500(w)
		return
	}
	if !isFriend && (userIdProfile != user.ID) {
		fmt.Println("isNotAccess - GetFollowersHandler error:", err)
		lib.HttpError400(w, "this user is not your friend")
		return
	}

	result, err := db.GetUserById(userIdProfile)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(" GetUserById - GetFollowersHandler error:", err)
			lib.HttpError404(w, "user not found ")
			return
		}
		fmt.Println("GetUserById - GetFollowersHandler error:", err)
		lib.HttpError500(w)
		return
	}

	resultUsers, err := db.GetFollowers(result)
	if err != nil {
		fmt.Println("EditProfile - GetFollowersHandler error:", err)
		lib.HttpError500(w)
		return
	}
	models.RemovePasswordOfUsers(resultUsers)
	jsonBytes, err := lib.ConvertToJsonBytes(resultUsers)
	if err != nil {
		fmt.Println("json.Marshal - GetFollowersHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)
}
