package middleware

import (
	"net/http"
	"sync"
	"time"
)

var (
	requests = make(map[string]int)
	mu       sync.Mutex
)

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		ip := r.RemoteAddr
		if _, exists := requests[ip]; !exists {
			requests[ip] = 0
		}

		requests[ip]++
		if requests[ip] > 1000 {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		go func() {
			time.Sleep(time.Hour)
			mu.Lock()
			defer mu.Unlock()
			requests[ip]--
		}()

		next.ServeHTTP(w, r)
	})
}
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow specific domains
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
