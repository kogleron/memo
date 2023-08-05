package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BotAPI interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}
