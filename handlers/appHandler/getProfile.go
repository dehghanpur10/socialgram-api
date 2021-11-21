package appHandler

import (
	"fmt"
	"net/http"
	"socialgram/lib"
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

	isFriend, err := db.IsFriend(user, userIdProfile)
	if err != nil {
		fmt.Println("IsFriend - GetProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}

	if !isFriend&&(userIdProfile!=user.ID){
		fmt.Println("this user is not friend - GetProfileHandler error:", err)
		lib.HttpError400(w, "this user is not friend")
		return
	}


	resultUser, err := db.GetProfileWithUserId(userIdProfile)
	if err != nil {
		fmt.Println("GetFriendsPosts - GetDashboardHandler error:", err)
		lib.HttpError500(w)
		return
	}

	resultUser.Password = ""

	jsonBytes, err := lib.ConvertToJsonBytes(resultUser)
	if err != nil {
		fmt.Println("json.Marshal - GetDashboardHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)

}


