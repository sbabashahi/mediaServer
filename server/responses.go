package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"time"
)

// Response struct
type Response struct {
	Data interface{} `json:"data"`
	Message string `json:"message"`
	Index int `json:"index"`
	Total int `json:"total"`
	CurrentTime int64 `json:"currentTime"`
	Status bool `json:"status"`
}

// JSONResponse Constructor
func JSONResponse(w http.ResponseWriter, data interface{}, message string, index int, total int, statusCode int) http.ResponseWriter {
	status := false
	if statusCode < 300 {
		status = true
	}
	nowTime := time.Now()
	w.WriteHeader(statusCode)
	resp :=  Response{Data: data, Message: message, Index: index, Total: total, CurrentTime: nowTime.Unix(), Status: status}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
	if status {
		log.Println(fmt.Sprintf("Success response status %v, Message: %v", statusCode, message))
	} else {
		log.Println(fmt.Sprintf("Error response status %v, Message: %v", statusCode, message))
	}
	return w
}