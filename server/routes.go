package server

import (
	"fmt"
	"net/http"

	"codeberg.org/ymazdy/mediamanager/media"
	"github.com/julienschmidt/httprouter"
)

//TODO: remove index test after initial phase
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

//TODO: implement connection to image save and resize
func uploadImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		parseErrorResponse(w)
		return
	}
	defer file.Close()
	uid := r.PostForm.Get("uid")
	// contentType := handler.Header.Get("Content-Type")

	media.SaveFileFromIOReader(uid+"/image.jpg", file)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello"))
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
