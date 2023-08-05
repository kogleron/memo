package command

import (
	"fmt"
	"strconv"
	"strings"

	"memo/internal/domain"
	"memo/internal/pkg/telegram"
)

func NewSearchExecutor(
	memoRepo domain.MemoRepository,
	resultsQty uint,
	replier telegram.Replier,
) *SearchExecutor {
	return &SearchExecutor{
		replier:    replier,
		memoRepo:   memoRepo,
		resultsQty: resultsQty,
	}
}

type SearchExecutor struct {
	replier    telegram.Replier
	memoRepo   domain.MemoRepository
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
	if cmd.Sender == nil {
		return errNeedStartCmd
	}

	text := strings.Trim(cmd.Payload, " ")
	if len(text) == 0 {
		return errEmptyPayload
	}

	memos, err := e.memoRepo.Search(text, cmd.Sender, e.resultsQty)
	if err != nil {
		return err
	}

	resp := e.getResponse(memos)

	return e.replier.ReplyTo(cmd.Message, resp)
}

func (e *SearchExecutor) getResponse(memos []domain.Memo) string {
	if len(memos) == 0 {
		return "nothing found"
	}

	text := "Found memos:"

	for _, memo := range memos {
		text += fmt.Sprintf("\n\n<b>#%d</b>: %s", memo.ID, memo.Text)
	}

	return text
}
