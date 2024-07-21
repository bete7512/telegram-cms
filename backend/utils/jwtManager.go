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

func GenerateSignupToken(user models.User) (string, error) {
	tokenClaims := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	return token.SignedString([]byte(config.JWT_SECRET))
}

func ValidateJwtToken(token string) (models.User, error) {
	tokenClaims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Id:    int(tokenClaims["id"].(float64)),
		Email: tokenClaims["email"].(string),
	}, nil
}

func ValidateToken(token string) (map[string]interface{}, error) {
	tokenClaims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	id := tokenClaims["id"].(float64)

	return map[string]interface{}{"id": id}, nil
}


func GenerateForgetPasswordToken(user models.User) (string, error) {
	tokenClaims := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 3).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	return token.SignedString([]byte(config.JWT_SECRET))
}