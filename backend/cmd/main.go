package main

import (
	"fmt"
	"log"

	"github.com/bete7512/telegram-cms/config"
	"github.com/bete7512/telegram-cms/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	host := config.DB_HOST
	port := config.DB_PORT
	user := config.DB_USER
	password := config.DB_PASSWORD
	dbname := config.DB_NAME

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	router := routes.Router(*db)
	if err := router.Run(fmt.Sprintf(":%s", config.PORT)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
