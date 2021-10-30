package lib

import "os"

var (
	SERVER_PORT string
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_ENGINE   string
)

func init() {
	SERVER_PORT = os.Getenv("SERVER_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_ENGINE = os.Getenv("DB_ENGINE")
}
