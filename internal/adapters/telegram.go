package adapters

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	bot          *tgbotapi.BotAPI
	openaiClient *OpenAIClient
}

func NewTelegramBot(token string, openaiClient *OpenAIClient) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TelegramBot{bot: bot, openaiClient: openaiClient}, nil
}

func (t *TelegramBot) isMentioned(msg *tgbotapi.Message) bool {
	for _, entity := range msg.Entities {
		if entity.Type == "mention" {
			mentionText := msg.Text[entity.Offset : entity.Offset+entity.Length]
			if mentionText == "@"+t.bot.Self.UserName {
				return true
			}
		}
	}
	return false
}

func (t *TelegramBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Chat.IsPrivate() || t.isMentioned(update.Message) {
				response, err := t.openaiClient.GetResponse(update.Message.Text)
				if err != nil {
					log.Printf("Error getting response from OpenAI: %v", err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
				t.bot.Send(msg)
			}
		}
	}
}
