package requestHandler

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"socialgram/lib"
)

func SetStatusRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - SetStatusRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - SetStatusRequestHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}

	friendId, err := lib.GetUserIdFromQuery(r)
	if err != nil {
		fmt.Println("GetUserIdFromQuery - SetStatusRequestHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}
	status, err := lib.GetStatusFromQuery(r)
	if err != nil {
		fmt.Println("GetStatusFromQuery - SetStatusRequestHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}
	friend, err := db.GetUserById(friendId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(" GetUserById - SetStatusRequestHandler error:", err)
			lib.HttpError404(w, "user not found ")
			return
		}
		fmt.Println("GetUserById - SetStatusRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}
	isRequest, err := db.IsRequest(friend, user.ID)
	if err != nil {
		fmt.Println("IsRequest - SetStatusRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}
	if !isRequest {
		fmt.Println("request already exist - SetStatusRequestHandler")
		lib.HttpError400(w, "request already not exist")
		return
	}
	if status == "ACCEPT" {
		err = db.CreateFriend(user, friendId)
		if err != nil {
			fmt.Println("DeleteFriend - SetStatusRequestHandler error:", err)
			lib.HttpError500(w)
			return
		}
	}
	err = db.DeleteRequest(friend, user.ID)
	if err != nil {
		fmt.Println("DeleteFriend - SetStatusRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}
	lib.HttpSuccessResponse(w, http.StatusNoContent, nil)
}
