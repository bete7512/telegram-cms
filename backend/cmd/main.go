package main

import (
	"github.com/bete7512/telegram-cms/routes"
)

func main() {
	router := routes.Router()
	router.Run(":8087")
}
