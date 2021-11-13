package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func RemoveLineBreakTrailing(data []byte) []byte {
	return []byte(strings.TrimSuffix(string(data), "\n"))
}
func InitLog(r *http.Request) {
	fmt.Printf("%s-%s-%v-", r.Method, r.URL.Path, r.URL.Query())
}

func HttpOptionsResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "X-PINGOTHER, Content-Type, authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, PUT, PATCH, GET, OPTIONS")
}

func jsonResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
}

func HttpSuccessResponse(w http.ResponseWriter, statusCode int, payload []byte) {
	jsonResponseHeaders(w)
	w.WriteHeader(statusCode)
	if statusCode == http.StatusNoContent {
		return
	}
	_, err := w.Write(payload)
	if err != nil {
		fmt.Print("HttpSuccessResponse-Write Error:", err)
		HttpError500(w)
	}
}

func HttpErrorResponse(w http.ResponseWriter, statusCode int, title, description string) {
	jsonResponseHeaders(w)
	w.WriteHeader(statusCode)
	errorResponse := NewErrorResponse(statusCode, title, description)
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		fmt.Print("HttpErrorResponse-json.Encode Error:", err)
		HttpError500(w)
	}
}

func HttpError400(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusBadRequest, "Bad request", description)
}
func HttpError401(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusUnauthorized, "Unauthorized", description)
}
func HttpError404(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusNotFound, "Not found", description)
}
func HttpError422(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusUnprocessableEntity, "unprocessable input", description)
}

func HttpError500(w http.ResponseWriter) {
	HttpErrorResponse(w, http.StatusInternalServerError, "Internal error", "Internal server error.")
}
