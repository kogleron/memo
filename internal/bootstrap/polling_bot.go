package bootstrap

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/apps"
	"memo/internal/command"
	"memo/internal/configs"
)

func NewPollingBot(tgBot *tgbotapi.BotAPI, tgConfig configs.TelegramConfig, cmdParser *command.Parser, cmdExecutors commandExecutors) *apps.PollingBot {
	return apps.NewPollingBot(tgBot, tgConfig, cmdParser, cmdExecutors, true)
}
