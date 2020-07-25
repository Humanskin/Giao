package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "qqa")
		c.Next()
		cost := time.Since(start)
		log.Println(cost)

	}
}