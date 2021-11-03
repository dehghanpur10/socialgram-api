package friendsHanlder

import (
	"fmt"
	"net/http"
	"socialgram/lib"
)

func DeleteFollowingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - DeleteFollowingHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - DeleteFollowingHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}
	friendId, err := lib.GetUserIdFromQuery(r)
	if err != nil {
		fmt.Println("GetUserIdFromQuery - DeleteFollowingHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	isFriend, err := db.IsFriend(user, friendId)
	if err != nil {
		fmt.Println("IsFriend - DeleteFollowingHandler error:", err)
		lib.HttpError500(w)
		return
	}

	if !isFriend {
		fmt.Println("is not friend - DeleteFollowingHandler error:", err)
		lib.HttpError400(w, "this user isn't your friend")
		return
	}

	err = db.DeleteFriend(user, friendId)
	if err != nil {
		fmt.Println("DeleteFriend - DeleteFollowingHandler error:", err)
		lib.HttpError500(w)
		return
	}
	lib.HttpSuccessResponse(w, http.StatusNoContent, nil)
}
