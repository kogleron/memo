package command

import (
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

func NewDeleteExecutor(
	memoRepo memo.Repository,
	userRepo user.Repository,
	replier telegram.Replier,
) *DeleteExecutor {
	return &DeleteExecutor{
		memoRepo: memoRepo,
		userRepo: userRepo,
		replier:  replier,
	}
}

type DeleteExecutor struct {
	memoRepo memo.Repository
	userRepo user.Repository
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
	text := strings.Trim(cmd.Payload, " ")
	if len(text) == 0 {
		return errEmptyPayload
	}

	id, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return err
	}

	user, err := e.getUser(cmd.Message)
	if err != nil {
		return err
	}

	memo, err := e.memoRepo.FindByID(user, uint(id))
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

func (e *DeleteExecutor) getUser(message *tgbotapi.Message) (*user.User, error) {
	if message == nil {
		return nil, telegram.ErrEmptyMessage
	}

	if message.From == nil {
		return nil, errEmptySender
	}

	user, err := e.userRepo.FindByTgAccount(message.From.UserName)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errNeedStartCmd
	}

	return user, nil
}
