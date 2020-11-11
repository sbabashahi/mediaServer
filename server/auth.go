package server

import (
	"net/http"
)

// AuthenticationMiddleware checks user authentication status
//TODO: proper authentication of the user
type AuthenticationMiddleware struct {
	Handler http.Handler
}

func (am AuthenticationMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if hasAccess(r) {
		am.Handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func hasAccess(r *http.Request) bool {
	return true
}
