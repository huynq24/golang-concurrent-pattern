package rate_limit

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
)

func Worker(id int, limiter *rate.Limiter) {
	for i := 0; i < 5; i++ {
		err := limiter.Wait(context.Background())
		if err != nil {
			return
		} // Wait for permission to proceed
		fmt.Printf("Worker %d: Job %d\n", id, i)
	}
}
