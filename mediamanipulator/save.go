package mediamanipulator

import (
	"io"
	"os"
)

func saveFileFromIOReader(path string, body io.ReadCloser) error {
	os.MkdirAll(path, 0777)
	fullPath := path + filename
	f, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, body)

	return err
}
