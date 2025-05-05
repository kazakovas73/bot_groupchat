package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	OpenAIToken   string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		OpenAIToken:   os.Getenv("OPENAI_API_KEY"),
	}

	if cfg.TelegramToken == "" || cfg.OpenAIToken == "" {
		log.Fatal("Missing required environment variables")
	}

	return cfg
}
