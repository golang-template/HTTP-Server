package middleware

import (
	"log"
	"net/http"
	"time"
)

// LogRequest is a simple middleware to log incoming requests
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// send request to next handler
		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("Completed %s in %v", r.URL.Path, duration)
	})
}
