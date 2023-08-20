package bootstrap

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/configs"
	"memo/internal/api/telegram"
	"memo/internal/api/telegram/command"
)

func NewTgAPIPollingBot(
	tgBot *tgbotapi.BotAPI,
	tgConfig configs.TelegramConfig,
	cmdParser command.Parser,
	cmdExecutors command.Executors,
	appConf configs.AppConfig,
) telegram.PollingBot {
	return telegram.NewPollingBot(
		tgBot,
		tgConfig,
		cmdParser,
		cmdExecutors,
		appConf.PollingShutdown,
	)
}
