package requestHandler

import (
	"net/http"
	"socialgram/lib"
)

func SetStatusRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
}