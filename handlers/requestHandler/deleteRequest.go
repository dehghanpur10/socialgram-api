package requestHandler

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"socialgram/lib"
)

func DeleteRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - DeleteRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - DeleteRequestHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}

	friendId, err := lib.GetUserIdFromQuery(r)
	if err != nil {
		fmt.Println("GetUserIdFromQuery - DeleteRequestHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}
	_, err = db.GetUserById(friendId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(" GetUserById - DeleteRequestHandler error:", err)
			lib.HttpError404(w, "user not found ")
			return
		}
		fmt.Println("GetUserById - DeleteRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}
	isRequest, err := db.IsRequest(user, friendId)
	if err != nil {
		fmt.Println("IsRequest - DeleteRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}
	if !isRequest {
		fmt.Println("request already exist - DeleteRequestHandler")
		lib.HttpError400(w, "request already not exist")
		return
	}
	err = db.DeleteRequest(user, friendId)
	if err != nil {
		fmt.Println("DeleteFriend - DeleteRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}
	lib.HttpSuccessResponse(w, http.StatusNoContent, nil)

}
