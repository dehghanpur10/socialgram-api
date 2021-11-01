package models

import (
	"time"
)


func NewLoginResponse(token string) *LoginResponse {
	response := new(LoginResponse)
	response.AccessToken = token
	response.EXP = time.Now().Add(time.Minute * 60 * 24 * 7).Unix()
	return response
}



