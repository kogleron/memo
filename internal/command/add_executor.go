package command

import (
	"memo/internal/memo"
)

type AddExecutor struct {
	memoRepo *memo.Repository
}

func (e *AddExecutor) Supports(cmd Command) bool {
	return cmd.Name == "add"
}

func (e *AddExecutor) Run(cmd Command) error {
	memo := &memo.Memo{
		Text: cmd.Payload,
	}

	e.memoRepo.Create(memo)

	return nil
}

func NewAddExecutor(memoRepo *memo.Repository) *AddExecutor {
	return &AddExecutor{
		memoRepo: memoRepo,
	}
}
