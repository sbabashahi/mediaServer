package main

import (
	"net/http"

	"codeberg.org/ymazdy/mediamanager/server"
)

func main() {
	router := server.GetRouter()

	nRouter := &server.AuthenticationMiddleware{Handler: router}

	http.ListenAndServe(":8080", server.FormParseMiddleware(nRouter))
}
