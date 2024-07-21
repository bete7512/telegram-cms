package utils

import "fmt"

var (
	ErrUserNotFound = fmt.Errorf("user not found")

	ErrUserAlreadyActive = fmt.Errorf("user already active")

	ErrUserAlreadyExists = fmt.Errorf("user already exists")

	ErrUserNotActive = fmt.Errorf("user not active")

	ErrWrongPassword = fmt.Errorf("wrong password")

	ErrInvalidToken = fmt.Errorf("invalid token")

	ErrTokenExpired = fmt.Errorf("token expired")

	ErrTokenMalformed = fmt.Errorf("token malformed")

	ErrTokenNotValidYet = fmt.Errorf("token not valid yet")

	ErrTokenInvalid = fmt.Errorf("token invalid")
)
