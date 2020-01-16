package ratelimiting_middleware

import (
	"fmt"
	"sync"
	"time"
)

//CounterRateLimitingProvider a simple rate limiting provider
type CounterRateLimitingProvider struct {
	counter         float64
	RequestMaxLimit float64
	requestsPerMS   float64
	mux             *sync.Mutex
}

//NewRateLimitingProvider returns Ratelimiting provider
func NewRateLimitingProvider(reqPerSec float64, reqWindow float64) RateLimiter {
	s := &CounterRateLimitingProvider{
		counter:         reqWindow,
		RequestMaxLimit: reqWindow,
		requestsPerMS:   reqPerSec,
		mux:             &sync.Mutex{},
	}

	go refreshCounter(s, reqPerSec, s.mux)
	return s
}

//RateLimit fetches data
func (crl *CounterRateLimitingProvider) RateLimit() bool {
	//check limit
	crl.mux.Lock()
	defer crl.mux.Unlock()

	if crl.counter < 1 {
		return false
	}

	crl.counter = crl.counter - 1
	return true
}

func refreshCounter(rl *CounterRateLimitingProvider, reqPerSec float64, m *sync.Mutex) {
	for now := range time.Tick(time.Second) {
		m.Lock()
		if rl.counter < rl.RequestMaxLimit {
			rl.counter = rl.counter + reqPerSec
		}
		m.Unlock()
		fmt.Printf("new counter: %v, req per ms: %v, req max limit: %v, time: %v\n", rl.counter, reqPerSec, rl.RequestMaxLimit, now)
	}
}
