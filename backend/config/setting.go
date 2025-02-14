package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT               string
	JWT_SECRET         string
	DB_URI             string
	ENVIROMENT         string
	SENDGRID_USER_NAME string
	SENDER_NAME        string
	DB_HOST            string
	DB_PORT            string
	DB_USER            string
	DB_PASSWORD        string
	DB_NAME            string
	SENDGRID_API_KEY   string
)

func init() {
	godotenv.Load()
	PORT = os.Getenv("PORT")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	DB_URI = os.Getenv("DB_URI")
	ENVIROMENT = os.Getenv("ENVIROMENT")
	SENDGRID_USER_NAME = os.Getenv("SENDGRID_USER_NAME")
	SENDER_NAME = os.Getenv("SENDER_NAME")
	SENDGRID_API_KEY = os.Getenv("SENDGRID_API_KEY")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	
}
