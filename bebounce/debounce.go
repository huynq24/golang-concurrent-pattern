package bebounce

import (
	"sync"
	"time"
)

func Debounce(f func(string), d time.Duration) func(string) {
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

		timer = time.AfterFunc(d, func() {
			mu.Lock()
			defer mu.Unlock()

			f(query)
		})
	}
}
