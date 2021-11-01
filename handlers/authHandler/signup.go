package authHandler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"socialgram/lib"
	"strings"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	db, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("GetDatabase - SignUpHandler error:", err)
		lib.HttpError500(w)
		return
	}

	userInput, err := lib.ParsUserInputFrom(r)
	if err != nil {
		fmt.Println("parseUserInput - SignUpHandler error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(userInput)
	if err != nil {
		fmt.Println(" validate - SignUpHandler error: ", err)
		lib.HttpError400(w, err.Error())
		return
	}

	if !lib.VerifyPassword(userInput.Password) {
		fmt.Println(" validate password - SignUpHandler  ")
		lib.HttpError400(w, "password should be correct")
		return
	}

	userInput.Password, err = lib.HashPassword(userInput.Password)
	if err != nil {
		fmt.Println(" HashPassword - SignUpHandler error: ", err)
		lib.HttpError500(w)
		return
	}

	userInput.AvatarURL, err = lib.SaveImageInStaticDirectory(r)
	if err != nil {
		if strings.Contains(err.Error(), "type") {
			fmt.Println(" saveImage - SignUpHandler error: ", err)
			lib.HttpError400(w, "image type is incorrect")
			return
		}
		if strings.Contains(err.Error(), "size") {
			fmt.Println(" saveImage - SignUpHandler error: ", err)
			lib.HttpError400(w, "image size is incorrect")
			return
		}
		fmt.Println("SaveImage - SignUpHandler - ", err)
		lib.HttpError500(w)
		return
	}

	err = db.CreateNewUser(userInput)
	if err != nil {
		removeErr := lib.RemoveImage(userInput.AvatarURL)
		if removeErr != nil {
			fmt.Println("remove image- CreateNewUser- SignUpHandler - ", err)
			lib.HttpError500(w)
			return
		}
		if strings.Contains(err.Error(), "Duplicate") {
			fmt.Println(" Duplicate - SignUpHandler error: ", err)
			lib.HttpError400(w, "username or email should be unique")
			return
		}
		fmt.Println(" CreateNewUser- SignUpHandler - ", err)
		lib.HttpError500(w)
		return
	}
	userInput.Password = ""
	jsonBytes, err := lib.ConvertToJsonBytes(userInput)
	if err != nil {
		fmt.Println("json.Marshal - SignUpHandler error:", err)
		lib.HttpError500(w)
		return
	}
	lib.HttpSuccessResponse(w, http.StatusCreated, jsonBytes)
}
