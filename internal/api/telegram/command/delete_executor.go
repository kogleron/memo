package command

import (
	"strconv"
	"strings"

	"memo/internal/domain"
	"memo/internal/pkg/telegram"
)

func NewDeleteExecutor(
	memoRepo domain.MemoRepository,
	replier telegram.Replier,
) *DeleteExecutor {
	return &DeleteExecutor{
		memoRepo: memoRepo,
		replier:  replier,
	}
}

type DeleteExecutor struct {
	memoRepo domain.MemoRepository
	replier  telegram.Replier
}

func (e DeleteExecutor) GetName() string {
	return "delete"
}

func (e DeleteExecutor) GetDescription() string {
	return "deletes the memo with a given id. \nFormat: /" + e.GetName() + " [MEMO_ID]"
}

func (e *DeleteExecutor) Supports(cmd Command) bool {
	return cmd.Name == e.GetName()
}

func (e *DeleteExecutor) Run(cmd Command) error {
	if cmd.Sender == nil {
		return errNeedStartCmd
	}

	text := strings.Trim(cmd.Payload, " ")
	if len(text) == 0 {
		return errEmptyPayload
	}

	id, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return err
	}

	memo, err := e.memoRepo.FindByID(cmd.Sender, uint(id))
	if err != nil {
		return err
	}

	if memo == nil {
		return e.replier.ReplyTo(cmd.Message, "not found")
	}

	err = e.memoRepo.Delete(memo)
	if err != nil {
		return err
	}

	return e.replier.ReplyTo(cmd.Message, "deleted")
}
