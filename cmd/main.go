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
		log.Fatalf("OpenAI init failed: %v", err)
	}

	telegramBot, err := adapters.NewTelegramBot(cfg.TelegramToken, openaiClient)
	if err != nil {
		log.Fatalf("Telegram bot init failed: %v", err)
	}

	telegramBot.Start()
}
