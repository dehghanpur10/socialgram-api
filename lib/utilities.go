package lib

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"unicode"
)

func SaveImageInStaticDirectory(r *http.Request) (string, error) {
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Print("read file - ")
		return "", err
	}
	defer file.Close()
	buff := make([]byte, 5*1024)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Print(" buffer file - ")
		return "", err
	}
	filetype := http.DetectContentType(buff)
	if !(filetype == "image/jpeg" || filetype == "image/jpg" || filetype == "image/png") {
		fmt.Print(" type incorrect - ")
		return "", errors.New("image type is incorrect")
	}
	if handler.Size > 5*1024*1024 {
		fmt.Print(" size incorrect - ")
		return "", errors.New("image size is incorrect")
	}
	newUUID, err := uuid.NewRandom()
	if err != nil {
		log.Println("SetUUID - ")
		return "", err
	}
	path := "/static/" + newUUID.String() + "-" + handler.Filename
	f, err := os.OpenFile("."+path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print(" openFile - ")
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Print(" copy file - ")
		return "", err
	}
	return path, nil
}

func RemoveImage(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Print("." + path)
	err = os.Remove(wd + path)
	if err != nil {
		return err
	}
	return nil
}

func VerifyPassword(s string) bool {
	var sevenOrMore, number, upper, letter, special bool
	letters := 0
	for _, c := range s {
		letters++
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letter = true
		default:
			//return false, false, false, false
		}
	}
	sevenOrMore = letters >= 8
	return sevenOrMore && number && upper && special && letter
}



func HashPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passwordBytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func CreateToken(username,email string) (string, error) {
	//create map claims
	claims := jwt.MapClaims{}
	//set data
	claims["username"] = username
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 60 * 24 * 7).Unix()
	//creat token without signed
	T := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	//signed token
	token, err := T.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", errors.New("error in signed")
	}
	return token, nil
}