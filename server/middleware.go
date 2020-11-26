package server

import (
    "context"
    "fmt"
	"net/http"
)

type customString string
const user customString = "user"
// User struct
type User struct {
    username string
    firstName string
    lastName string
    id int
}
//MyMiddleware is good
func MyMiddleware(handler http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {
        ctx := context.WithValue(r.Context(), user, User{"saeed70", "saeed", "babashahi", 127})
		handler.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(middleware)
}


// FormParseMiddleware parses user form data
func FormParseMiddleware(handler http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {
		if err:=r.ParseMultipartForm(10 << 20); err != nil {
			JSONResponse(w, nil, fmt.Sprint(err), 0, 0, 400)
			return
		}
		err := r.ParseForm()
		if err != nil {
			JSONResponse(w, nil, fmt.Sprint(err), 0, 0, 400)
			return
		}
		fmt.Println(r.Header.Get("Authorization"))
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(middleware)
}
