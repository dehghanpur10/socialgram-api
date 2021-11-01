package models

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	EXP         int64 `json:"exp"`
}
