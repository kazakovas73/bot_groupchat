package main

import (
	"bot_groupchat/internal/adapters"
	"bot_groupchat/internal/config"
	"log"
)

func main() {
	cfg := config.Load()

	openaiClient, err := adapters.NewOpenAIClient(cfg.OpenAIToken)
	if err != nil {
		log.Fatalf("Не удалось создать OpenAI клиента: %v", err)
	}

	bot, err := adapters.NewTelegramBot(cfg.TelegramToken, openaiClient)
	if err != nil {
		log.Fatalf("Не удалось запустить бота: %v", err)
	}

	bot.Start()
}
