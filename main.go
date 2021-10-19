package main

import (
	"fmt"
	"net/http"
	"socialgram/lib"
	"socialgram/routes"
)

func main() {
	fmt.Println("server run ...")

	router := routes.Init()
	http.Handle("/api/v1/", router)

	err := http.ListenAndServe(lib.SERVER_PORT, router)
	if err != nil {
		fmt.Println("main - problem in run server - error: ", err)
	}
}
