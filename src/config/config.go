package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	return os.Getenv("TELEGRAM_BOT_TOKEN")
}
