package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"socialgram/handlers/httpserver"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(httpserver.NotFoundController)

	return r
}