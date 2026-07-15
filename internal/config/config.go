package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL     string
	CoinGeckoURL    string
	CoinGeckoAPIKey string
	TelegramAPIKey  string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("no .env file found")
	}
	return &Config{
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		CoinGeckoURL:    os.Getenv("COINGECKO_API_URL"),
		CoinGeckoAPIKey: os.Getenv("COINGECKO_API_KEY"),
		TelegramAPIKey:  os.Getenv("TELEGRAM_API_KEY"),
	}
}
