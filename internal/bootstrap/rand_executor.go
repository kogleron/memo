package bootstrap

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/command"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/user"
)

func NewRandExecutor(
	conf configs.AppConfig,
	memoRepo *memo.Repository,
	tgBot *tgbotapi.BotAPI,
	userRepo *user.Repository,
) *command.RandExecutor {
	return command.NewRandExecutor(uint(conf.RandQty), memoRepo, tgBot, userRepo)
}
