package main

import (
	"go-concurrent-parttern/rate_limit"
	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(rate.Limit(2), 1)
	for i := 0; i < 3; i++ {
		go rate_limit.Worker(i, limiter)
	}

	//time.Sleep(10 * time.Second)
	//router := gin.Default()
	//handleDebounced := bebounce.Debounce(func(query string) {
	//	fmt.Printf("Performing search for: %s\n", query)
	//}, 5*time.Second)

	//limiter := rate.NewLimiter(rate.Every(time.Second), 1)

	//handleDebounced := rate_limit_debounce.DebounceWithRateLimit(func(query string) {
	//	fmt.Printf("Performing search for: %s\n", query)
	//}, 2*time.Second, limiter)

	// Define an API endpoint to handle input events
	//router.GET("/search", func(c *gin.Context) {
	//	query := c.Query("query")
	//	handleDebounced(query)
	//	c.JSON(http.StatusOK, gin.H{"message": "Search request received"})
	//})

	//router.Run(":8080")
}
