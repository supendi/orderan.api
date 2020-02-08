package security

import (
	"net/http"
	"strings"
)

//JWTAuthMiddleware is jwt middleware
type JWTAuthMiddleware struct {
	jwtKey       string
	tokenHandler TokenHandler
}

//NewJWTAuthMiddleware return a new JWTAuthMiddleware
func NewJWTAuthMiddleware(jwtKey string, tokenHandler TokenHandler) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{
		jwtKey:       jwtKey,
		tokenHandler: tokenHandler,
	}
}

//HandlerFunc return handler func as middleware to check auth token from request header
func (me *JWTAuthMiddleware) HandlerFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		jwtToken := strings.ReplaceAll(authToken, "Bearer ", "")
		if !me.tokenHandler.Verify(jwtToken, me.jwtKey) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
