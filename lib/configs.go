package lib

import "os"

var(
	SERVER_PORT string
)

func init() {
	SERVER_PORT = os.Getenv("SERVER_PORT")
}