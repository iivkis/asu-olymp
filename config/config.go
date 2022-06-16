package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	//server settings
	PORT string

	//other
	JWT_SECRET string
)

func init() {
	PORT = os.Getenv("PORT")
	JWT_SECRET = os.Getenv("JWT_SECRET")
}
