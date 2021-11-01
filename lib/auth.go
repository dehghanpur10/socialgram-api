package lib

import (
	"errors"
	"fmt"
	"net/http"
	"socialgram/models"
	"strings"
)

func GetBearerUser(db SocialGramStore, headers http.Header) (*models.User, error) {
	bearerToken, err := GetHeaderValue(headers, "Authorization")
	if err != nil {
		fmt.Print("GetBearerUser-")
		return nil, err
	}

	token, err := newBearerUser(bearerToken)
	if err != nil {
		fmt.Print("GetBearerUser-bearer:{", bearerToken, "}-")
		return nil, err
	}

	ok, userData, err := ExtractToken(token)

	if err != nil || !ok {
		fmt.Print("token is not verify")
		return nil, errors.New("token is not verify")
	}

	user, err := db.GetUser(userData["username"].(string))
	if err != nil {
		fmt.Print("user not found")
		return nil, errors.New("user of this token not found")
	}

	return user, nil
}

func GetHeaderValue(headers http.Header, headerKey string) (string, error) {
	headerValue, exists := headers[headerKey]
	if !exists {
		err := errors.New("Couldn't find " + headerKey + " in request headers.")
		fmt.Print("GetHeaderValue-")
		return "", err
	}
	if len(headerValue) == 0 {
		err := errors.New("Empty value " + headerKey + " in request headers.")
		fmt.Print("GetHeaderValue-")
		return "", err
	}
	return headerValue[0], nil
}

func newBearerUser(bearerToken string) (string, error) {
	if len(bearerToken) == 0 || !strings.HasPrefix(bearerToken, "Bearer ") {
		fmt.Print("newBearerUser-")
		return "", errors.New("Invalid bearer token.")
	}
	splitValue := strings.Split(bearerToken, " ")
	if len(splitValue) != 2 {
		fmt.Print("newBearerUser-")
		return "", errors.New("The bearer token has missing parts.")
	}

	return splitValue[1], nil
}
