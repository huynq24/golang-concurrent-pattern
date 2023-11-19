package main

import (
	"fmt"
	"go-concurrent-parttern/bebounce"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	time.Sleep(5 * time.Second)
	router := gin.Default()

	handleDebounced := bebounce.Debounce(func(query string) {
		fmt.Printf("Performing search for: %s\n", query)
	}, 5*time.Second)

	// limiter := rate.NewLimiter(rate.Every(time.Second), 2)
	// handleDebounced := rate_limit_debounce.DebounceWithRateLimit(func(query string) {
	// 	fmt.Printf("Performing search for: %s\n", query)
	// }, 5*time.Second, limiter)

	// Define an API endpoint to handle input events
	router.GET("/search", func(c *gin.Context) {
		query := c.Query("query")
		handleDebounced(query)
		c.JSON(http.StatusOK, gin.H{"message": "Search request received"})
	})

	router.Run(":8080")
}
