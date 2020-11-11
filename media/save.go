package media

import (
	"io"
	"os"
	"strings"
)

// TODO: needs efficient method
func folderFromFilePath(path string) string {
	parts := strings.Split(path, "/")
	folder := ""
	if len(parts) > 1 {
		for i := 0; i < len(parts)-1; i++ {
			folder += parts[i] + "/"
		}
	}
	if []rune(path)[0] == '/' {
		folder = "/" + folder
	}
	return folder
}

// SaveFileFromIOReader saves a file from io reader to a path
func SaveFileFromIOReader(path string, body io.ReadCloser) error {
	os.MkdirAll(folderFromFilePath(path), 0777)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, body)

	return err
}
