package media

import (
	"fmt"
	"strings"
	"time"
)

var basePath = "./"

// PathMaker makes a path based on section (company, user) and uid (or company slug)
func PathMaker(section string, uid string) string {
	path := fmt.Sprintf("%s%s/", basePath, section)
	if uid != "" {
		path = fmt.Sprintf("%s%s/", path, uid)
	}
	return path
}

// NameMaker makes a file name based on time and content type of file
func NameMaker(contentType string) string {
	contentTypeParts := strings.Split(contentType, "/")
	now := time.Now()
	filename := fmt.Sprintf("%s_%d_%d_%d_%d_%d.%s", contentTypeParts[0], int(now.Year()), now.Month(), now.Day(), now.Hour(), now.Nanosecond(), contentTypeParts[1])
	return filename
}
