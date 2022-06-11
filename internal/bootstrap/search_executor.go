package bootstrap

import (
	"memo/internal/command"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

func NewSearchExecutor(tgBot telegram.BotAPI, memoRepo memo.Repository, userRepo user.Repository, conf configs.AppConfig) *command.SearchExecutor {
	return command.NewSearchExecutor(tgBot, memoRepo, userRepo, conf.SearchResultQty)
}
