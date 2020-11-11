package media

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

var basePath = ""

//TODO: get sizes from config map instead of hardcoding them
func generateThumb(filename string, path string) {
	// createThumb(356, 200, filename, path)
	// createThumb(383, 300, filename, path)
	createThumb(80, 80, filename, path)
	createThumb(56, 56, filename, path)
	createThumb(44, 44, filename, path)
}

func createThumb(width int, height int, fileName string, path string) {
	newPath := basePath + strconv.Itoa(width) + "x" + strconv.Itoa(height) + "/" + strings.Replace(path, basePath, "", -1)
	if _, err := os.Stat(newPath + fileName); os.IsNotExist(err) {
		src, err := imaging.Open(path + fileName)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		dstImage := imaging.Resize(src, width, height, imaging.Lanczos)
		os.MkdirAll(newPath, 0777)
		imaging.Save(dstImage, newPath+fileName)
	}

}
