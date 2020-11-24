package server

import (
	"fmt"
	"net/http"

	"codeberg.org/ymazdy/mediamanager/media"
	"github.com/julienschmidt/httprouter"
)

//TODO: remove index test after initial phase
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w = JSONResponse(w, make(map[string]int, 0), "Welcome", 0, 0, 200)
}

//MARK: upload user image route
//TODO: implement connection to image save and resize
func uploadImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// e := govalidator.New(govalidator.Options{
	// 	Request:         r,
	// 	Rules:           imageUploadValidator,
	// 	RequiredDefault: true,
	// }).Validate()
	// if e != nil {
	// 	err := map[string]interface{}{"validationError": e}
	// 	w.Header().Set("Content-type", "application/json")
	// 	json.NewEncoder(w).Encode(err)
	// }

	file, fileHeader, err := r.FormFile("uploadfile")
	if err != nil {
		parseErrorResponse(w)
		return
	}
	defer file.Close()
	uid := r.PostForm.Get("uid")
	contentType := fileHeader.Header.Get("Content-Type")
	
	if err = checkFileContentType(contentType);err != nil {
		JSONResponse(w, nil, fmt.Sprint(err), 0, 0, 400)
		return
	}
	if err = checkFileSize(fileHeader.Size);err != nil {
		JSONResponse(w, nil, fmt.Sprint(err), 0, 0, 400)
		return
	}
	path := media.PathMaker("user", uid)
	name := media.NameMaker(contentType)

	media.SaveFileFromIOReader(path, name, file)

	JSONResponse(w, ImageResponse{path+name}, "Upload Success", 0, 0, 200)
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
