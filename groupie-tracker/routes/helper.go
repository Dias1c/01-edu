package routes

import "net/http"

// Set headers for CORS Policy
func setCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin")
}

// UnUSING!!!
// Middleware Set Cors
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin")

		next.ServeHTTP(rw, r)
	})
}
