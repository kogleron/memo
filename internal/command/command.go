package command

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Command struct {
	Name    string
	Payload string
	Message *tgbotapi.Message
}
