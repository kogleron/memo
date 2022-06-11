package bootstrap

import (
	"memo/internal/apps"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

func NewRandCommand(
	conf configs.AppConfig,
	memoRepo memo.Repository,
	userRepo user.Repository,
	tgBot telegram.BotAPI,
) *apps.RandCommand {
	return apps.NewRandCommand(uint(conf.RandQty), memoRepo, userRepo, tgBot)
}
