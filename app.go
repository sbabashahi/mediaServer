package main

import (
	"fmt"
	"log"
	"net/http"

	"codeberg.org/ymazdy/mediamanager/server"
)

func main() {
	router := server.GetRouter()

	// nRouter := &server.AuthenticationMiddleware{Handler: router}
	nRouter := server.FormParseMiddleware(router)

	fmt.Println("Starting the Server")
	log.Fatal(http.ListenAndServe(":8080", nRouter))
}
