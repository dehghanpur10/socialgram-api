package lib

import (
	"os"
	"strconv"
)

var (
	SERVER_PORT string
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_ENGINE   string
	SECRET_KEY  string
	PAGE_SIZE   int
)

func init() {
	SERVER_PORT = os.Getenv("PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_ENGINE = os.Getenv("DB_ENGINE")
	SECRET_KEY = os.Getenv("SECRET_KEY")
	PAGE_SIZE, _ = strconv.Atoi(os.Getenv("PAGE_SIZE"))
}
