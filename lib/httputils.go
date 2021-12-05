package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"socialgram/models"
	"strconv"
)

func ConvertToJsonBytes(payload interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(payload)
	return buffer.Bytes(), err
}

//func ParsUserInputFrom(r *http.Request) (*models.User, error) {
//	userInput := new(models.User)
//	userInput.Username = r.FormValue("username")
//	userInput.Name = r.FormValue("name")
//	userInput.Email = r.FormValue("email")
//	userInput.Password = r.FormValue("password")
//	userInput.Gender = r.FormValue("gender")
//	age, err := strconv.ParseUint(r.FormValue("age"), 10, 32)
//	userInput.Age = uint(age)
//	userInput.City = r.FormValue("city")
//	userInput.Country = r.FormValue("country")
//	userInput.Bio = r.FormValue("bio")
//	userInput.Interest = r.FormValue("interest")
//	_, _, err = r.FormFile("image")
//	return userInput, err
//}

func ParseUserInputFrom(jsonBytes []byte) (*models.User, error) {
	user := new(models.User)
	err := json.Unmarshal(jsonBytes, user)
	return user, err
}

func GetUserInfoFromPath(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	if id, exists := vars["userInfo"]; exists {
		return id, nil
	}
	return "", errors.New("invalid userInfo in path")
}

func GetPageNumberFromQuery(r *http.Request) (int, error) {
	values, exists := r.URL.Query()["page"]
	if !exists || len(values) == 0 || len(values[0]) == 0 {
		return 0, errors.New("page not found in query")
	}
	parseInt, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, errors.New("invalid page in query. Must be a number")
	}

	return parseInt, nil
}

func GetPostIdFromQuery(r *http.Request) (uint, error) {
	values, exists := r.URL.Query()["post_id"]
	if !exists || len(values) == 0 || len(values[0]) == 0 {
		return 0, errors.New("post_id not found in query")
	}
	parseInt, err := strconv.ParseUint(values[0], 10, 32)
	if err != nil {
		return 0, errors.New("invalid post_id in query. Must be a number")
	}

	return uint(parseInt), nil
}

func ParsPostInputFrom(jsonBytes []byte) (*models.Post, error) {
	postInput := new(models.Post)
	err := json.Unmarshal(jsonBytes, postInput)
	return postInput, err
}

func GetUserIdFromQuery(r *http.Request) (uint, error) {
	values, exists := r.URL.Query()["user_id"]
	if !exists || len(values) == 0 || len(values[0]) == 0 {
		return 0, errors.New("user_id not found in query")
	}
	parseInt, err := strconv.ParseUint(values[0], 10, 32)
	if err != nil {
		return 0, errors.New("invalid user_id in query. Must be a number")
	}

	return uint(parseInt), nil
}
func GetStatusFromQuery(r *http.Request) (string, error) {
	values, exists := r.URL.Query()["status"]
	if !exists || len(values) == 0 || len(values[0]) == 0 {
		return "", errors.New("status not found in query")
	}

	if values[0] == "ACCEPT" || values[0] == "DECLINE" {
		return values[0], nil
	}
	return "", errors.New("invalid status")
}
