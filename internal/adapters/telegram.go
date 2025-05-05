package adapters

import (
	"context"
	"log"
	"runtime"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	bot          *tgbotapi.BotAPI
	openaiClient *OpenAIClient
	jobs         chan MessageJob
}

type MessageJob struct {
	Message *tgbotapi.Message
}

func NewTelegramBot(token string, openaiClient *OpenAIClient) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TelegramBot{
		bot:          bot,
		openaiClient: openaiClient,
		jobs:         make(chan MessageJob, 100),
	}, nil
}

func (t *TelegramBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.bot.GetUpdatesChan(u)

	numCpu := runtime.NumCPU()
	log.Printf("Number of CPU: %d", numCpu)

	for range numCpu / 2 {
		go t.worker()
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Chat.IsPrivate() || t.isMentioned(update.Message) {
			t.jobs <- MessageJob{Message: update.Message}
		}
	}
}

func (t *TelegramBot) worker() {
	for job := range t.jobs {
		t.handleMessage(job.Message)
	}
}

func (t *TelegramBot) handleMessage(msg *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := t.openaiClient.GetResponse(ctx, msg.Text)
	if err != nil {
		log.Printf("OpenAI error: %v", err)
		t.sendAndLogMessage(msg.Chat.ID, defaultAnswer)
	} else {
		t.sendAndLogMessage(msg.Chat.ID, response)
	}
}

func (t *TelegramBot) sendAndLogMessage(chatID int64, text string) {
	messageConfig := tgbotapi.NewMessage(chatID, text)
	_, err := t.bot.Send(messageConfig)
	if err != nil {
		log.Printf("sendMessage error: %v", err)
		return
	}
	// log.Printf("Sent to chat %d: %s", msg.Chat.ID, msg.Text)
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
