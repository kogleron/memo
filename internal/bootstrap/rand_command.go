package bootstrap

import (
	"memo/configs"
	"memo/internal/apps"
	"memo/internal/domain"
	"memo/internal/pkg/telegram"
)

func NewRandCommand(
	conf configs.AppConfig,
	memoRepo domain.MemoRepository,
	userRepo domain.UserRepository,
	tgBot telegram.BotAPI,
) *apps.RandCommand {
	return apps.NewRandCommand(uint(conf.RandQty), memoRepo, userRepo, tgBot)
}
