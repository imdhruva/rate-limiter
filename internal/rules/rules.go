package rules

import (
	"rate-limiter/internal/fixed_token"
)

var clientBucketMap = make(map[string]*fixed_token.TokenBucket)

type Rule struct {
	Rate      float64
	MaxTokens float64
}

// GetBucket returns the bucket for a specific user type
func GetBucket(identifier string, usertype string) *fixed_token.TokenBucket {
	// if client does not exist then assign one
	if clientBucketMap[identifier] == nil {
		clientBucketMap[identifier] = fixed_token.NewTokenBucket(rulesMap[usertype].MaxTokens, rulesMap[usertype].Rate)
	}
	return clientBucketMap[identifier]
}

var rulesMap = map[string]Rule{
	"default": {Rate: 5, MaxTokens: 5},
}
