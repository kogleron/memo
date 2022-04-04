package command

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/memo"
)

type AddExecutor struct {
	memoRepo *memo.Repository
	tgBot    *tgbotapi.BotAPI
}

func (e *AddExecutor) Supports(cmd Command) bool {
	return cmd.Name == "add"
}

func (e *AddExecutor) Run(cmd Command) error {
	memo := &memo.Memo{
		Text: cmd.Payload,
	}

	e.memoRepo.Create(memo)

	if cmd.Message == nil {
		return nil
	}

	msg := tgbotapi.NewMessage(
		cmd.Message.Chat.ID,
		"done",
	)
	msg.ReplyToMessageID = cmd.Message.MessageID

	_, err := e.tgBot.Send(msg)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func NewAddExecutor(memoRepo *memo.Repository, tgBot *tgbotapi.BotAPI) *AddExecutor {
	return &AddExecutor{
		memoRepo: memoRepo,
		tgBot:    tgBot,
	}
}
