package fixed_token

import (
	"math"
	"sync"
	"time"
)

// TokenBucket holds token details
type TokenBucket struct {
	tokens         float64
	limit          float64
	rate           float64
	lastAccessTime time.Time
	mu             sync.Mutex
}

// NewTokenBucket initiates tokenBucket object
func NewTokenBucket(limit, rate float64) *TokenBucket {
	return &TokenBucket{
		tokens:         limit,
		limit:          limit,
		rate:           rate,
		lastAccessTime: time.Now(),
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	duration := now.Sub(tb.lastAccessTime)
	tokenToAdd := tb.rate * duration.Seconds()
	tb.tokens = math.Min(tokenToAdd+tb.tokens, tb.limit)
	tb.lastAccessTime = now
}

// Request implements token bucket for a given token
func (tb *TokenBucket) Request(tokens float64) (bool, float64) {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()
	if tokens <= tb.tokens {
		tb.tokens -= tokens
		return true, tb.tokens
	}
	return false, 0
}
