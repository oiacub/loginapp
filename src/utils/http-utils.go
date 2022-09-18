package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

func CookieExist(w http.ResponseWriter, r *http.Request) bool {
	err := r.ParseForm()
	_, err = r.Cookie("token")
	if err != nil {
		return false
	}
	return true
}

func GetToken(r *http.Request) string {
	return r.Header.Get("Token")
}

func GetGoogleToken(r *http.Request) string {
	return r.Header.Get("GoogleToken")
}

func GenerateCookie(w http.ResponseWriter, token string) {

	expires := time.Now().AddDate(1, 0, 0)
	ck := http.Cookie{
		Name:    "token",
		Domain:  "localhost:8080",
		Path:    "/",
		Expires: expires,
	}
	ck.Value = token
	http.SetCookie(w, &ck)
}

func SendResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func SendResponseJSON(w http.ResponseWriter, message []byte, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(message)
}
