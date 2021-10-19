package postHandler

import (
	"net/http"
	"socialgram/lib"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
}