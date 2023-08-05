package bootstrap

import (
	"memo/configs"
	"memo/internal/api/telegram/command"
	"memo/internal/domain"
	"memo/internal/pkg/telegram"
)

func NewSearchExecutor(
	memoRepo domain.MemoRepository,
	conf configs.AppConfig,
	replier telegram.Replier,
) *command.SearchExecutor {
	return command.NewSearchExecutor(memoRepo, conf.SearchResultQty, replier)
}
