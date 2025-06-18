package ratelimit

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/hvmello/goal-tracker-backend/internal/config"
	"github.com/hvmello/goal-tracker-backend/internal/goals"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	config   config.RateLimitConfig
}

func New(config config.RateLimitConfig) *RateLimiter {
	log.Printf("Initializing rate limiter: enabled=%v, requests per minute=%d, burst size=%d",
		config.Enabled, config.RequestsPerMinute, config.BurstSize)

	return &RateLimiter{
		visitors: make(map[string]*visitor),
		config:   config,
	}
}

func (rl *RateLimiter) getVisitor(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	if !exists {
		// Use rate.Every to define the interval between requests
		//If 5 req/min, we de 1 req every 12 seconds
		interval := time.Minute / time.Duration(rl.config.RequestsPerMinute)
		limiter := rate.NewLimiter(rate.Every(interval), rl.config.BurstSize)

		log.Printf("New visitor: %s, interval: %v, burst: %d",
			ip, interval, rl.config.BurstSize)

		rl.visitors[ip] = &visitor{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

func (rl *RateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	for ip, v := range rl.visitors {
		if time.Since(v.lastSeen) > 3*time.Hour {
			delete(rl.visitors, ip)
			log.Printf("Cleaned up visitor: %s", ip)
		}
	}
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	log.Println("Middleware started for rate limiting")

	go func() {
		for {
			time.Sleep(time.Hour)
			rl.cleanup()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.config.Enabled || strings.HasPrefix(r.URL.Path, "/swagger/") {
			next.ServeHTTP(w, r)
			return
		}

		// Gets IP for client
		ip := r.Header.Get("X-Real-IP")
		if ip == "" {
			ip = r.Header.Get("X-Forwarded-For")
		}
		if ip == "" {
			ip = r.RemoteAddr
		}

		// Limiter for IP
		limiter := rl.getVisitor(ip)

		// 429 if too many requests
		if !limiter.Allow() {
			log.Printf("Rate limit exceeded for IP: %s, path: %s", ip, r.URL.Path)

			w.Header().Set("Retry-After", "60")
			w.Header().Set("X-RateLimit-Limit", "5")
			w.Header().Set("X-RateLimit-Remaining", "0")

			goals.WriteErrorResponse(w, goals.ErrTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
