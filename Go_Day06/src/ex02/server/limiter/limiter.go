package limmiter

import (
	// "log"
	"net/http"
	"strings"
	"sync"
)

type RateLimiter struct {
	requests map[string]int
	mu       sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		requests: make(map[string]int),
	}
}

func (rl *RateLimiter) LimiterMidlewire(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientIP := strings.Split(r.RemoteAddr, ":")
		rl.mu.Lock()
		defer rl.mu.Unlock()

		rl.requests[clientIP[0]]++
		if rl.requests[clientIP[0]] > 100 {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}
