package models

import (
	"time"
)



func NewLoginResponse(token string) *ResponseLogin {
	response := new(ResponseLogin)
	response.AccessToken = token
	response.EXP = time.Now().Add(time.Minute * 60 * 24 * 7).Unix()
	return response
}

func RemovePasswordOfUsers(users []User)  {
	for i := 0; i < len(users); i++ {
		users[i].Password = ""
	}
}
