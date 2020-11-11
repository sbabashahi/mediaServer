package server

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/julienschmidt/httprouter"
)

var decoder = schema.NewDecoder()

type imageUploadRequest struct {
	uid        string
	uploadfile multipart.File
}

//TODO: remove index test after initial phase
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

//TODO: implement connection to image save and resize
func uploadImage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req imageUploadRequest
	err := decoder.Decode(&req, r.PostForm)
	if err != nil {
		parseErrorResponse(w)
	}
	defer req.uploadfile.Close()
}

// GetRouter returns the default server router
func GetRouter() *httprouter.Router {
	router := httprouter.New()

	//MARK: implementing cors headers for frontend
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}
		w.WriteHeader(http.StatusNoContent)
	})

	router.GET("/", index)
	router.POST("/image/", uploadImage)

	return router
}
