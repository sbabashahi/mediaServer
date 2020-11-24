package server

import (
	"encoding/json"
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"codeberg.org/ymazdy/mediamanager/media"
	"github.com/nfnt/resize"
)

// FormParseMiddleware parses user form data
func FormParseMiddleware(handler http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {
		if err:=r.ParseMultipartForm(10 << 20); err != nil {
			JSONResponse(w, nil, fmt.Sprint(err), 0, 0, 400)
		}// 10 MB
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(400)
		}
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(middleware)
}

func parseErrorResponse(w http.ResponseWriter) {
	fmt.Fprint(w, "Error parsing the request!\n")
}

func jsonResponse(w http.ResponseWriter, v interface{}) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if v != nil {
		body, err := json.Marshal(v)
		if err != nil {
			return err
		}
		// can just return an error when connection is hijacked or content-size is longer then declared.
		if _, err := w.Write(body); err != nil {
			return err
		}
	}

	return nil
}

func checkFileContentType(contentType string) error {
	supportedMediaTypes := map[string]bool{"image/jpeg": true, "image/png": true}
	_, exist := supportedMediaTypes[contentType]
	if !exist {
		message := fmt.Sprintf("Not supported content- %v ", contentType)
		return fmt.Errorf(message)
	}
	return nil
}

func checkFileSize(size int64) error {
	maxFileSize := int64(10*1024*1024)  // 10 MB
	if size > maxFileSize {
		message := fmt.Sprintf("Max file size is %d but your file size is %d", maxFileSize, size)
		return fmt.Errorf(message)
	}
	return nil
}

func resizeImage(filePath string) error {
	
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	// resize file to this width and height
	imageSize := map[int]int{100: 100, 200: 200, 300: 300}
	for width, height := range imageSize {
		m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
		out, err := os.Create(media.ResizeNameMaker(filePath, width, height))
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
	
		// write new image to file
		jpeg.Encode(out, m, nil)
	}
	
	return nil
}