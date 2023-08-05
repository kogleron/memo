package bootstrap

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/configs"
)

func NewTgBot(tgConf configs.TelegramConfig) (*tgbotapi.BotAPI, error) {
	tgBot, err := tgbotapi.NewBotAPI(tgConf.BotToken)
	if err != nil {
		return nil, err
	}

	tgBot.Debug = tgConf.DebugMode

	return tgBot, nil
}
