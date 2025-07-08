package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost  string `env:"DB_HOST"`
	DBPort  string `env:"DB_PORT"`
	DBUser  string `env:"DB_USER"`
	DBPass  string `env:"DB_PASS"`
	DBName  string `env:"DB_NAME"`
	SSLMode string `env:"SSL_MODE"`
}

func Load() *Config {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg := &Config{
		DBHost:  os.Getenv("DB_HOST"),
		DBPort:  os.Getenv("DB_PORT"),
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASS"),
		DBName:  os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("SSL_MODE"),
	}
	return cfg
}
