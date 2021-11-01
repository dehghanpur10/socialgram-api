package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	uuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"os"
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

func ConvertToJsonBytes(payload interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(payload)
	return buffer.Bytes(), err
}

func HashPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passwordBytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}