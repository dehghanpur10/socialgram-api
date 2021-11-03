package friendsHanlder

import (
	"fmt"
	"net/http"
	"socialgram/lib"
)

func GetFollowingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - GetFollowingHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - GetFollowingHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}

	friends, err := db.GetFriends(user)
	if err != nil {
		fmt.Println("GetFriends - GetFollowingHandler error:", err)
		lib.HttpError500(w)
		return
	}

	jsonBytes, err := lib.ConvertToJsonBytes(friends)
	if err != nil {
		fmt.Println("json.Marshal - GetFollowingHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)

}
