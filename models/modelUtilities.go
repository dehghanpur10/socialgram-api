package models

import (
	"net/http"
	"strconv"
)

func ParsUserInputFrom(r *http.Request) (*User, error) {
	userInput := new(User)
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


