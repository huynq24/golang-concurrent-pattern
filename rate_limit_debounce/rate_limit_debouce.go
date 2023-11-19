package rate_limit_debounce

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func DebounceWithRateLimit(f func(string), d time.Duration, r *rate.Limiter) func(string) {
	var timer *time.Timer
	var query string
	var mu sync.Mutex

	return func(q string) {
		mu.Lock()
		defer mu.Unlock()

		query = q

		if timer != nil {
			timer.Stop()
		}

		// Use the rate limiter to control the rate of requests
		if r.Allow() {
			timer = time.AfterFunc(d, func() {
				mu.Lock()
				defer mu.Unlock()

				f(query)
			})
		} else {
			fmt.Println("Rate limited")
		}
	}
}
