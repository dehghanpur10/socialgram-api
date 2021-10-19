package requestHandler

import (
	"net/http"
	"socialgram/lib"
)

func CreateRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
}