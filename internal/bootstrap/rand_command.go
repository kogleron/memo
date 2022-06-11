package bootstrap

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/apps"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/user"
)

func NewRandCommand(
	conf configs.AppConfig,
	memoRepo *memo.Repository,
	userRepo *user.Repository,
	tgBot *tgbotapi.BotAPI,
) *apps.RandCommand {
	return apps.NewRandCommand(uint(conf.RandQty), memoRepo, userRepo, tgBot)
}
