package requestHandler

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"socialgram/lib"
)

func CreateRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - CreateRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - CreateRequestHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}

	friendId, err := lib.GetUserIdFromQuery(r)
	if err != nil {
		fmt.Println("GetUserIdFromQuery - CreateRequestHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	isFriend, err := db.IsFriend(user, friendId)
	if err != nil {
		fmt.Println("IsFriend - CreateRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}
	if isFriend || (friendId == user.ID) {
		fmt.Println("IsFriend - CreateRequestHandler error:", err)
		lib.HttpError400(w, "you can not send request")
		return
	}

	_, err = db.GetUserById(friendId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(" GetUserById - CreateRequestHandler error:", err)
			lib.HttpError404(w, "user not found ")
			return
		}
		fmt.Println("GetUserById - CreateRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}
	isRequest, err := db.IsRequest(user, friendId)
	if err != nil {
		fmt.Println("IsRequest - CreateRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}
	if isRequest {
		fmt.Println("request already exist - CreateRequestHandler")
		lib.HttpError400(w, "request already exist")
		return
	}
	err = db.CreateRequest(user, friendId)
	if err != nil {
		fmt.Println("CreateRequest - CreateRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusNoContent, nil)
}
