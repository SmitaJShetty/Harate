package ratelimiting_middleware

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_RateLimit_OK(t *testing.T) {
	requestPerSecond := 0.1
	reqWindow := 60.0
	rl := NewRateLimitingProvider(requestPerSecond, reqWindow)
	res := rl.RateLimit()
	assert.Equal(t, true, res)
}

func Test_RateLimit_TooManyRequest(t *testing.T) {
	requestsPerSecond := 0.00000001
	reqWindow := 0.1
	rl := NewRateLimitingProvider(requestsPerSecond, reqWindow)
	res := rl.RateLimit()
	assert.Equal(t, false,res)
}
