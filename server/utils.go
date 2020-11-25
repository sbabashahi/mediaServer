package server

import (
	"encoding/json"
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jtguibas/cinema"
	"github.com/nfnt/resize"

	"codeberg.org/ymazdy/mediamanager/media"
)

var supportedMediaTypes map[string]int64

func init() {
    supportedMediaTypes = map[string]int64{
		"image/jpeg": int64(10*1024*1024),
		"image/png": int64(10*1024*1024),
		"video/mp4": int64(50*1024*1024),
		}
}

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
	_, exist := supportedMediaTypes[contentType]
	if !exist {
		message := fmt.Sprintf("Not supported content- %v ", contentType)
		return fmt.Errorf(message)
	}
	return nil
}

func checkFileSize(size int64, contentType string) error {
	maxFileSize := supportedMediaTypes[contentType]
	if size > maxFileSize {
		message := fmt.Sprintf("Max file size for %s is %d but your file size is %d", contentType, maxFileSize, size)
		return fmt.Errorf(message)
	}
	return nil
}

func mediaConvertor(filePath, contentType string) error {
	
	if strings.HasPrefix(contentType, "image/") {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
			return err
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
	} else if strings.HasPrefix(contentType, "video"){
		video, err := cinema.Load(filePath)
		if err != nil {
			log.Fatal(err)
			return err
		}
		width, height := video.Width()/10, video.Height()/10
		video.SetSize(width, height)
		path := media.ResizeNameMaker(filePath, width, height)
		video.Render(path)  // check for another way of converting, it increse audio and video bitrate :|
	}
	
	return nil
}