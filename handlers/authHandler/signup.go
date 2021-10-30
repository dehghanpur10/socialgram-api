package authHandler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
	"socialgram/lib"
	"socialgram/models"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}

	lib.InitLog(r)

	_, err := lib.GetDatabase()
	if err != nil {
		fmt.Println("SignUpHandler-GetDatabase error:", err)
		lib.HttpError500(w)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("SignUpHandler-ioutil.ReadAll error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	userInput, err := models.ParsUserInputFrom(reqBody)
	if err != nil {
		fmt.Println("SignUpHandler-ParseWorkbenchInputFrom error:", err)
		lib.HttpError400(w, err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(userInput)
	if err != nil {
		fmt.Println("SignUpHandler error: ",err)
		lib.HttpError400(w, err.Error())
		return
	}

}
