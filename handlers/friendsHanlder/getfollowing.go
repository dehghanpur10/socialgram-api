package friendsHanlder

import (
	"net/http"
	"socialgram/lib"
)

func GetFollowingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		lib.HttpOptionsResponseHeaders(w)
		return
	}
}
