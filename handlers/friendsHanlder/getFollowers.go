package friendsHanlder

import (
	"fmt"
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

	resultUsers, err := db.GetFollowers(user)
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
