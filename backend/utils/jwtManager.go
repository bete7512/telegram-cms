package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/bete7512/telegram-cms/config"
	"github.com/bete7512/telegram-cms/models"
)

func GenerateJWT(user models.User) (string, error) {
	tokenClaims := jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	return token.SignedString([]byte(config.JWT_SECRET))
}
