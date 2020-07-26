package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		time.Sleep(3 * time.Millisecond)
		c.Set("name", "qqa")
		c.Next()
		cost := time.Since(start)
		log.Printf("time is %v", cost)

	}
}