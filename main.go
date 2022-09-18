package main

import (
	apiauth "gotestweb/src/routes/api/v1/auth"
	apiusers "gotestweb/src/routes/api/v1/users"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/", fileServer))

	apiusers.RegisterRequests()
	apiauth.RegisterRequests()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
