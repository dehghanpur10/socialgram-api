package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"socialgram/lib"
	"socialgram/routes"
)

func main() {
	fmt.Println("server run ...")
	router := routes.Init()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "PATCH", "POST", "PUT","DELETE", "OPTIONS"})
	err := http.ListenAndServe(lib.SERVER_PORT, handlers.CORS(originsOk, headersOk, methodsOk)(router))
	if err != nil {
		fmt.Println("main - problem in run server - error: ", err)
	}
}
