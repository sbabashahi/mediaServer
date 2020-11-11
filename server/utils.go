package server

import (
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
