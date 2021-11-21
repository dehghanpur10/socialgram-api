package appHandler

import (
	"fmt"
	"net/http"
	"socialgram/lib"
	"socialgram/models"
)

func GetDashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - GetDashboardHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - GetDashboardHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}

	pageNumber, err := lib.GetPageNumberFromQuery(r)
	if err != nil {
		fmt.Println("GetPageNumberFromQuery - GetDashboardHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	posts, err := db.GetFriendsPosts(user, pageNumber)
	if err != nil {
		fmt.Println("GetFriendsPosts - GetDashboardHandler error:", err)
		lib.HttpError500(w)
		return
	}

	models.RemovePasswordFromPosts(posts)

	jsonBytes, err := lib.ConvertToJsonBytes(posts)
	if err != nil {
		fmt.Println("json.Marshal - GetDashboardHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)
}
