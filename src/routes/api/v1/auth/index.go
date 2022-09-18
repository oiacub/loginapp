package auth

import (
	"net/http"
)

func RegisterRequests() {
	http.HandleFunc("/auth/loginByGoogle", HandleGoogleLoginExisting)
	http.HandleFunc("/auth/createGoogleUser", HandleGoogleLogin)
	http.HandleFunc("/auth/dologin", LoginUser)
	http.HandleFunc("/auth/logout", HandleLogout)
}
