package users

import (
	"net/http"
)

func RegisterRequests() {
	http.HandleFunc("/users/get", Get)
	http.HandleFunc("/users/create", Create)
	http.HandleFunc("/users/update", UpdateProfile)
}
