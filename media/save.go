package media

import (
	"io"
	"os"
)

// SaveFileFromIOReader saves a file from io reader to a path
func SaveFileFromIOReader(path string, name string, body io.ReadCloser) error {
	os.MkdirAll(path, 0777)
	f, err := os.Create(path + name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, body)

	return err
}
