package utils

import (
    "log"
    "strings"
)

const (
    DUPLICATE_EMAIL_ERROR = "users_email_key"
)

func FilterError(err error) (int, string) {
    log.Println("Error:", err.Error())
    if strings.Contains(err.Error(), DUPLICATE_EMAIL_ERROR) {
        return 409, "This Email already exists"
    } 
    return 500, err.Error()
}
