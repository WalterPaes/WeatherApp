package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type vars struct {
	AccuweatherBaseUrl string
	AccuweatherApiKey  string
}

func Load() *vars {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &vars{
		AccuweatherBaseUrl: os.Getenv("ACCUWEATHER_BASE_URL"),
		AccuweatherApiKey:  os.Getenv("ACCUWEATHER_API_KEY"),
	}
}
