package requestHandler

import (
	"fmt"
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

	err = db.CreateRequest(user, friendId)
	if err != nil {
		fmt.Println("CreateRequest - CreateRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusNoContent, nil)
}