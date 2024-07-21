package middleware

import (
	"github.com/bete7512/telegram-cms/models"
	"github.com/bete7512/telegram-cms/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		tokenString := authHeader[len("Bearer "):]
		user, err := utils.ValidateJwtToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		if user == (models.User{}) {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
