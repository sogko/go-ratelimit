package ratelimit

//
// Ratelimiting incoming connections - Small Library
//
// (c) 2015 Hafiz Ismail (@sogko) <hafiz-at-wehavefaces-net>
//
// Notes:
// - Modified implementation by Sudhi Herle
// - Used per second unit time.
//   Previously it was erroneously in nanosecond
// - Allowed user to specify per unit time (for eg: 500 messages per 8 seconds)
//
//
// Credits:
// - Sudhi Herle <sudhi-dot-herle-at-gmail-com>
//   - Reference: https://code.google.com/p/go-wiki/wiki/RateLimiting
// - Anti Huimaa
//   - Reference: http://stackoverflow.com/questions/667508/whats-a-good-rate-limiting-algorithm
//
// License: GPLv2
//
// Original notes:
//  - This is a very simple interface for rate limiting. It
//    implements a token bucket algorithm
//  - Based on Anti Huimaa's very clever token bucket algorithm.
//
// Usage:
//    import "github.com/sogko/go-ratelimit"
//
//    // rate limit at 500 messages every 8 seconds
//    rate := 500
//    per := 8
//    rl := ratelimit.NewRateLimiter(rate, per)
//
//    ....
//    if rl.Limit() {
//       drop_connection(conn)
//    }
//
import (
	"time"
)

type Ratelimiter struct {
	rate int       // conn per unit
	per  int       // unit
	last time.Time // last time we were polled/asked

	allowance float64
}

// Create new rate limiter that limits at rate/sec
func NewRateLimiter(rate int, per int) (*Ratelimiter, error) {
	r := Ratelimiter{rate: rate, last: time.Now(), per: per}

	r.allowance = float64(r.rate)
	return &r, nil
}

// Return true if the current call exceeds the set rate, false otherwise
func (r *Ratelimiter) Limit() bool {

	// handle cases where rate in config file is unset - defaulting
	// to "0" (unlimited)
	if r.rate == 0 {
		return false
	}

	if r.per <= 0 {
		r.per = 1
	}

	rate := float64(r.rate)
	per := float64(r.per)
	now := time.Now()
	elapsed := now.Sub(r.last).Seconds()

	r.last = now
	r.allowance += float64(elapsed) * (rate / per)

	// Clamp number of tokens in the bucket. Don't let it get
	// unboundedly large
	if r.allowance > rate {
		r.allowance = rate
	}

	var ret bool

	if r.allowance < 1.0 {
		ret = true
	} else {
		r.allowance -= 1.0
		ret = false
	}

	return ret
}
