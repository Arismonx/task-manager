package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENV() {
	// Load Environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
