package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FormParseMiddleware parses user form data
func FormParseMiddleware(handler http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
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
