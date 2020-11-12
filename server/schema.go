package server

import "github.com/thedevsaddam/govalidator"

var imageUploadValidator = govalidator.MapData{
	"uploadfile": []string{"required"},
	"uid":        []string{"required", "min:2", "max:32"},
}

type ImageResponse struct {
	Message string `json:"message"`
	Path    string `json:"path"`
}
