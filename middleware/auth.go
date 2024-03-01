package middleware

import "net/http"

// AuthMiddleware is an example of a simple authentication middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authentication logic would go here
		next.ServeHTTP(w, r)
	})
}
