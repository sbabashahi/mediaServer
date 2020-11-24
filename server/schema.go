package server

import "github.com/thedevsaddam/govalidator"

var imageUploadValidator = govalidator.MapData{
	"uploadfile": []string{"required"},
	"uid":        []string{"required", "min:2", "max:32"},
}

// ImageResponse struct
type ImageResponse struct {
	Path    string `json:"path"`
}
