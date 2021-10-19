package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"socialgram/handlers/appHandler"
	"socialgram/handlers/authHandler"
	"socialgram/handlers/friendsHanlder"
	"socialgram/handlers/httpserver"
	"socialgram/handlers/postHandler"
	"socialgram/handlers/requestHandler"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", authHandler.SignUpHandler).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/login", authHandler.LoginHandler).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/search/{userInfo}", appHandler.SearchHandler).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/profile", appHandler.GetProfileHandler).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/profile", appHandler.EditProfileHandler).Methods(http.MethodPut, http.MethodOptions)

	r.HandleFunc("/dashboard/{page}", appHandler.GetDashboardHandler).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/post", postHandler.CreatePostHandler).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/post", postHandler.LikePostHandler).Methods(http.MethodPatch, http.MethodOptions)

	r.HandleFunc("/post", postHandler.DeletePostHandler).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/request", requestHandler.CreateRequestHandler).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/request", requestHandler.DeleteRequestHandler).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/request", requestHandler.GetRequestHandler).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/request", requestHandler.SetStatusRequestHandler).Methods(http.MethodPut, http.MethodOptions)

	r.HandleFunc("/followers", friendsHanlder.GetFollowersHandler).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/followers", friendsHanlder.DeleteFollowersHandler).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/following", friendsHanlder.GetFollowingHandler).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/following", friendsHanlder.DeleteFollowingHandler).Methods(http.MethodDelete, http.MethodOptions)

	r.NotFoundHandler = http.HandlerFunc(httpserver.NotFoundController)

	return r
}