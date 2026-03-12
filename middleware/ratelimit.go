package middleware

import (
    "net/http"
    "sync"
    "time"
    "github.com/gin-gonic/gin"
)

type RateLimiter struct {
    requests map[string][]time.Time
    mu       sync.RWMutex
    limit    int
    window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    return &RateLimiter{
        requests: make(map[string][]time.Time),
        limit:    limit,
        window:   window,
    }
}

func (rl *RateLimiter) isAllowed(ip string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    now := time.Now()
    var valid []time.Time
    for _, t := range rl.requests[ip] {
        if now.Sub(t) < rl.window {
            valid = append(valid, t)
        }
    }
    if len(valid) >= rl.limit {
        rl.requests[ip] = valid
        return false
    }
    rl.requests[ip] = append(valid, now)
    return true
}

func RateLimit() gin.HandlerFunc {
    limiter := NewRateLimiter(60, time.Minute)
    return func(c *gin.Context) {
        if !limiter.isAllowed(c.ClientIP()) {
            c.JSON(http.StatusTooManyRequests, gin.H{"code": 429, "msg": "Too many requests"})
            c.Abort()
            return
        }
        c.Next()
    }
}
