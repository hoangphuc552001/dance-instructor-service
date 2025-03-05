package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Auth AuthConfig
	HTTP HTTPConfig
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		panic(err)
	}

	return &Config{
		Auth: LoadAuthConfig(),
		HTTP: LoadHTTPConfig(),
	}
}
