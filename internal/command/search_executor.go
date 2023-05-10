package command

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

func NewSearchExecutor(
	memoRepo memo.Repository,
	userRepo user.Repository,
	resultsQty uint,
	replier telegram.Replier,
) *SearchExecutor {
	return &SearchExecutor{
		replier:    replier,
		memoRepo:   memoRepo,
		userRepo:   userRepo,
		resultsQty: resultsQty,
	}
}

type SearchExecutor struct {
	replier    telegram.Replier
	memoRepo   memo.Repository
	userRepo   user.Repository
	resultsQty uint
}

func (e SearchExecutor) GetName() string {
	return "search"
}

func (e SearchExecutor) GetDescription() string {
	return "retrivies the memos with a given text." +
		" Max qty of memos in result is " + strconv.FormatUint(uint64(e.resultsQty), 10) + "." +
		" \nFormat: " + e.GetName() + " [SEARCH_TEXT]"
}

func (e SearchExecutor) Supports(cmd Command) bool {
	return cmd.Name == e.GetName()
}

func (e *SearchExecutor) Run(cmd Command) error {
	text := strings.Trim(cmd.Payload, " ")
	if len(text) == 0 {
		return errEmptyPayload
	}

	user, err := e.getUser(cmd.Message)
	if err != nil {
		return err
	}

	memos, err := e.memoRepo.Search(text, user, e.resultsQty)
	if err != nil {
		return err
	}

	resp := e.getResponse(memos)

	return e.replier.ReplyTo(cmd.Message, resp)
}

func (e *SearchExecutor) getResponse(memos []memo.Memo) string {
	if len(memos) == 0 {
		return "nothing found"
	}

	text := "Found memos:"

	for _, memo := range memos {
		text += fmt.Sprintf("\n\n<b>#%d</b>: %s", memo.ID, memo.Text)
	}

	return text
}

func (e *SearchExecutor) getUser(message *tgbotapi.Message) (*user.User, error) {
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
