package command

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

type SearchExecutor struct {
	tgBot      telegram.BotAPI
	memoRepo   memo.Repository
	userRepo   user.Repository
	resultsQty uint
}

func (e *SearchExecutor) Supports(cmd Command) bool {
	return cmd.Name == "search"
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

	return e.replyTo(cmd.Message, resp)
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

func (e *SearchExecutor) replyTo(message *tgbotapi.Message, text string) error {
	if message == nil {
		return errEmptyMessage
	}

	if message.Chat == nil {
		return errNoChatID
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		text,
	)

	msg.ReplyToMessageID = message.MessageID
	msg.DisableWebPagePreview = true
	msg.ParseMode = tgbotapi.ModeHTML

	_, err := e.tgBot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (e *SearchExecutor) getUser(message *tgbotapi.Message) (*user.User, error) {
	if message == nil {
		return nil, errEmptyMessage
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

func NewSearchExecutor(tgBot telegram.BotAPI, memoRepo memo.Repository, userRepo user.Repository, resultsQty uint) *SearchExecutor {
	return &SearchExecutor{
		tgBot:      tgBot,
		memoRepo:   memoRepo,
		userRepo:   userRepo,
		resultsQty: resultsQty,
	}
}
