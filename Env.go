package main

import (
	"github.com/joho/godotenv"
	"log"
)

func env() {
	// Load environment variables.
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
