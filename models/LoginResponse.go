package models

type ResponseLogin struct {
	AccessToken string `json:"access_token"`
	EXP         int64  `json:"exp"`
}
