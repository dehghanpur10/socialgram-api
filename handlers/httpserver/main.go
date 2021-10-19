package httpserver

import (
	"net/http"
	"socialgram/lib"
)

func NotFoundController(w http.ResponseWriter, r *http.Request) {
	lib.HttpError404(w, "Requested resource doesn't exist. Please check your path.")
}
