package bootstrap

import (
	"memo/internal/command"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

func NewRandExecutor(
	conf configs.AppConfig,
	memoRepo memo.Repository,
	tgBot telegram.BotAPI,
	userRepo user.Repository,
) *command.RandExecutor {
	return command.NewRandExecutor(uint(conf.RandQty), memoRepo, tgBot, userRepo)
}
