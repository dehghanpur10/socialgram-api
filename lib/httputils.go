package lib

import (
	"bytes"
	"encoding/json"
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

func ParsUserInputFrom(r *http.Request) (*models.User, error) {
	userInput := new(models.User)
	userInput.Username = r.FormValue("username")
	userInput.Name = r.FormValue("name")
	userInput.Email = r.FormValue("email")
	userInput.Password = r.FormValue("password")
	userInput.Gender = r.FormValue("gender")
	age, err := strconv.ParseUint(r.FormValue("age"), 10, 32)
	userInput.Age = uint(age)
	userInput.City = r.FormValue("city")
	userInput.Country = r.FormValue("country")
	userInput.Bio = r.FormValue("bio")
	userInput.Interest = r.FormValue("interest")
	_, _, err = r.FormFile("image")
	return userInput, err
}

func ParseLoginUserInputFrom(jsonBytes []byte) (*models.User, error) {
	user := new(models.User)
	err := json.Unmarshal(jsonBytes, user)
	return user, err
}

