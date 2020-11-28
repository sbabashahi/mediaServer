package main

import (
	"fmt"
	"log"
	"net/http"

	"codeberg.org/ymazdy/mediamanager/server"
)

func main() {
	router := server.GetRouter()
	fmt.Println("Starting the Server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
