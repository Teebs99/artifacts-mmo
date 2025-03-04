package main

import (
	"github.com/Teebs99/artifacts-mmo/models"
	_ "github.com/joho/godotenv"
	"os"
)

func main() {
	config := models.Config{
		ApiKey: os.Getenv("API_KEY"),
	}

	client := getClient(&config)

}
