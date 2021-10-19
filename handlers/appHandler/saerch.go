package appHandler

import (
	"net/http"
	"socialgram/lib"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
}

