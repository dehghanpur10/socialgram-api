package requestHandler

import (
	"fmt"
	"net/http"
	"socialgram/lib"
)

func GetRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - GetRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - GetRequestHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}
	resultUsers, err := db.GetRequests(user)
	if err != nil {
		fmt.Println("EditProfile - GetRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}
	jsonBytes, err := lib.ConvertToJsonBytes(resultUsers)
	if err != nil {
		fmt.Println("json.Marshal - GetRequestHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)
}
