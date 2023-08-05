package bootstrap

import (
	"memo/configs"
	"memo/internal/api/telegram/command"
	"memo/internal/domain"
	"memo/internal/pkg/telegram"
)

func NewRandExecutor(
	conf configs.AppConfig,
	memoRepo domain.MemoRepository,
	tgBot telegram.BotAPI,
) *command.RandExecutor {
	return command.NewRandExecutor(uint(conf.RandQty), memoRepo, tgBot)
}
