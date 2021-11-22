package appHandler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"socialgram/lib"
)

func EditProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - EditProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}

	user, err := lib.GetBearerUser(db, r.Header)
	if err != nil {
		fmt.Println("GetBearerUser - EditProfileHandler error:", err)
		lib.HttpError401(w, err.Error())
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll  - EditProfileHandler error:", err)
		lib.HttpError400(w, "invalid request structure")
		return
	}

	userInput, err := lib.ParseLoginUserInputFrom(reqBody)
	if err != nil {
		fmt.Println("parseUserInput - EditProfileHandler error:", err)
		lib.HttpError400(w, "invalid user input")
		return
	}

	resultUser, err := db.EditProfile(user, userInput)
	if err != nil {
		fmt.Println("EditProfile - EditProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}

	resultUser.Password = ""

	jsonBytes, err := lib.ConvertToJsonBytes(resultUser)
	if err != nil {
		fmt.Println("json.Marshal - EditProfileHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)

}
