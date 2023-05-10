package bootstrap

import (
	"memo/internal/command"
	"memo/internal/configs"
	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

func NewSearchExecutor(
	memoRepo memo.Repository,
	userRepo user.Repository,
	conf configs.AppConfig,
	replier telegram.Replier,
) *command.SearchExecutor {
	return command.NewSearchExecutor(memoRepo, userRepo, conf.SearchResultQty, replier)
}
