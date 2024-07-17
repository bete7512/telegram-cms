package middleware

import "github.com/gin-gonic/gin"

func VerifyJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verify JWT
		c.Next()
	}
}
