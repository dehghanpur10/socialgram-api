package appHandler

import (
	"fmt"
	"net/http"
	"socialgram/lib"
	"socialgram/models"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - SearchHandler error:", err)
		lib.HttpError500(w)
		return
	}

	_, err = lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - SearchHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	userInfo, err := lib.GetUserInfoFromPath(r)
	if err != nil {
		fmt.Println("GetUserInfoFromPath - SearchHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	pageNumber, err := lib.GetPageNumberFromQuery(r)
	if err != nil {
		fmt.Println("GetPageNumberFromQuery - SearchHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	users, err := db.SearchUsers(userInfo, pageNumber)
	if err != nil {
		fmt.Println("SearchUsers - SearchHandler error:", err)
		lib.HttpError500(w)
		return
	}

	models.RemovePasswordOfUsers(users)

	jsonBytes, err := lib.ConvertToJsonBytes(users)
	if err != nil {
		fmt.Println("json.Marshal - SearchHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)

}
