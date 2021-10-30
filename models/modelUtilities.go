package models

import "encoding/json"

func ParsUserInputFrom(jsonBytes []byte) (*User, error) {
	userInput := new(User)
	err := json.Unmarshal(jsonBytes, userInput)
	return userInput, err
}
