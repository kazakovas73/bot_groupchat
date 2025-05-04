package main

import (
	"bot_groupchat/src/adapters"
	"bot_groupchat/src/config"
	"log"
)

func main() {
	token := config.LoadToken()

	bot, err := adapters.NewTelegramBot(token)
	if err != nil {
		log.Fatalf("Не удалось запустить бота: %v", err)
	}

	bot.Start()
}
