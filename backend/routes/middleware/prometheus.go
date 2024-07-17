package middleware

import "github.com/gin-gonic/gin"

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prometheus metrics
		c.Next()
	}
}