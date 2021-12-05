package postHandler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
	"socialgram/lib"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - CreatePostHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - CreatePostHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll  - SignUpHandler error:", err)
		lib.HttpError400(w, "invalid request structure")
		return
	}

	postInput, err := lib.ParsPostInputFrom(reqBody)
	if err != nil {
		fmt.Println("parsePostInput - CreatePostHandler error:", err)
		lib.HttpError400(w, "image field not found")
		return
	}

	validate := validator.New()
	err = validate.Struct(postInput)
	if err != nil {
		fmt.Println(" validate - CreatePostHandler error: ", err)
		lib.HttpError400(w, "invalid post input")
		return
	}

	postInput.UserID = user.ID

	err = db.CreateNewPost(postInput)
	if err != nil {
		fmt.Println(" CreateNewPost- CreatePostHandler - ", err)
		lib.HttpError500(w)
		return
	}
	user.Password = ""
	postInput.User = user

	jsonBytes, err := lib.ConvertToJsonBytes(postInput)
	if err != nil {
		fmt.Println("json.Marshal - CreatePostHandler error:", err)
		lib.HttpError500(w)
		return
	}
	lib.HttpSuccessResponse(w, http.StatusCreated, jsonBytes)
}
