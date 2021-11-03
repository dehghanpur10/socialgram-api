package postHandler

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"socialgram/lib"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - LikePostHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - LikePostHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}

	postId, err := lib.GetPostIdFromQuery(r)
	if err != nil {
		fmt.Println("GetPostIdFromQuery - LikePostHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	post, err := db.GetPost(postId)
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

	isFriend, err := db.IsFriend(user, post.UserID)
	if err != nil {
		fmt.Println("IsFriend - LikePostHandler error:", err)
		lib.HttpError500(w)
		return
	}

	if !isFriend && post.UserID != user.ID{
		fmt.Println("is not friend - LikePostHandler error:", err)
		lib.HttpError400(w,"you don't access to this post")
		return
	}

	status, err := db.GetLikeStatus(postId, user.ID)
	if err != nil {
		fmt.Println("GetLikeStatus - LikePostHandler error:", err)
		lib.HttpError500(w)
		return
	}

	status, err = db.ToggleLike(status, postId, user)
	if err != nil {
		fmt.Println("ToggleLike - LikePostHandler error:", err)
		lib.HttpError500(w)
		return
	}
	response := fmt.Sprintf(`{"status":%v}`, status)
	lib.HttpSuccessResponse(w, http.StatusOK, []byte(response))
}
