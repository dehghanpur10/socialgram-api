package authHandler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"socialgram/lib"
	"socialgram/models"
	"strings"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - LoginHandler error:", err)
		lib.HttpError500(w)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll  - LoginHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	userInput, err := lib.ParseLoginUserInputFrom(reqBody)
	if err != nil {
		fmt.Println("parseUserInput - LoginHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	if len(userInput.Password) == 0 || len(userInput.Username) == 0 {
		fmt.Println("empty password or username - LoginHandler error:", err)
		lib.HttpError400(w, "empty password or username")
		return
	}

	user, err := db.GetUser(userInput.Username)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			fmt.Println(" GetUser - LoginHandler error:", err)
			lib.HttpError400(w, "not found username")
			return
		}
		fmt.Println("GetUser - LoginHandler error:", err)
		lib.HttpError500(w)
		return
	}

	check := lib.CheckPasswordHash(userInput.Password, user.Password)
	if !check {
		fmt.Println("CheckPasswordHash - LoginHandler")
		lib.HttpError400(w, "password is incorrect")
		return
	}

	token, err := lib.CreateToken(user.Username, user.Email)
	if err != nil {
		fmt.Println("CreateToken - LoginHandler error:", err)
		lib.HttpError500(w)
		return
	}

	response := models.NewLoginResponse(token)

	jsonBytes, err := lib.ConvertToJsonBytes(response)
	if err != nil {
		fmt.Println("json.Marshal - LoginHandler error:", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, jsonBytes)

}
