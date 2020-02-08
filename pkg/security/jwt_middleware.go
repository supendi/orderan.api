package security

import (
	"net/http"
	"strings"
)

//NewJWTAuthMiddleware return handler func middleware to check auth token from request header
func NewJWTAuthMiddleware(tokenHandler TokenHandler, jwtKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("Authorization")
			if authToken == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			jwtToken := strings.ReplaceAll(authToken, "Bearer ", "")
			if !tokenHandler.Verify(jwtToken, jwtKey) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
