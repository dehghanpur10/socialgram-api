package postHandler

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"socialgram/lib"
	"strings"
)

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - DeletePostHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - DeletePostHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}

	postId, err := lib.GetPostIdFromQuery(r)
	if err != nil {
		fmt.Println("GetPostIdFromQuery - DeletePostHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	_, err = db.GetPost(postId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(" GetPost - LikePostHandler error:", err)
			lib.HttpError404(w, "post not found ")
			return
		}
		fmt.Println("GetPost - LikePostHandler error:", err)
		lib.HttpError500(w)
		return
	}

	err = db.DeletePost(postId,user.ID)
	fmt.Println(err)
	if err != nil {
		if strings.Contains(err.Error(), "access") {
			fmt.Println(" DeletePost - DeletePostHandler error:", err)
			lib.HttpError404(w, err.Error())
			return
		}
		fmt.Println("DeletePost - DeletePostHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusNoContent, nil)
}
