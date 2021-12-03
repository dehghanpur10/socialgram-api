package friendsHanlder

import (
	"fmt"
	"net/http"
	"socialgram/lib"
)

func DeleteFollowersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - DeleteFollowersHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - DeleteFollowersHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}
	friendId, err := lib.GetUserIdFromQuery(r)
	if err != nil {
		fmt.Println("GetUserIdFromQuery - DeleteFollowersHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}
	friend, err := db.GetUserById(friendId)
	if err != nil {
		fmt.Println("GetUserById - DeleteFollowersHandler error:", err)
		lib.HttpError400(w, "this user not found")
		return
	}
	isFriend, err := db.IsFriend(friend, user.ID)
	if err != nil {
		fmt.Println("IsFriend - DeleteFollowersHandler error:", err)
		lib.HttpError500(w)
		return
	}

	if !isFriend {
		fmt.Println("is not friend - DeleteFollowersHandler error:")
		lib.HttpError400(w, "this user isn't your follower")
		return
	}

	err = db.DeleteFriend(friend, user.ID)
	if err != nil {
		fmt.Println("DeleteFriend - DeleteFollowersHandler error:", err)
		lib.HttpError500(w)
		return
	}
	lib.HttpSuccessResponse(w, http.StatusNoContent, nil)
}
