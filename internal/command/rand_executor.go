package command

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/memo"
)

type RandExecutor struct {
	defaultMemoQty uint
	memoRepo       *memo.Repository
	tgBot          *tgbotapi.BotAPI
}

func (e *RandExecutor) Supports(cmd Command) bool {
	return cmd.Name == "rand"
}

func (e *RandExecutor) Run(cmd Command) error {
	if cmd.Message == nil {
		return errNoCmdMessage
	}

	message := cmd.Message

	memos := e.memoRepo.Rand(e.defaultMemoQty)
	if len(memos) == 0 {
		return e.onNoMemos(message)
	}

	return e.sendMemos(message, memos)
}

func (e *RandExecutor) sendMemos(message *tgbotapi.Message, memos []memo.Memo) error {
	var err error

	for i := range memos {
		msg := tgbotapi.NewMessage(
			message.Chat.ID,
			memos[i].Text,
		)
		msg.ReplyToMessageID = message.MessageID

		_, err = e.tgBot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}

	return err
}

func (e *RandExecutor) onNoMemos(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		"no memos",
	)
	msg.ReplyToMessageID = message.MessageID

	_, err := e.tgBot.Send(msg)

	return err
}

func NewRandExecutor(defaultMemoQty uint, memoRepo *memo.Repository, tgBot *tgbotapi.BotAPI) *RandExecutor {
	return &RandExecutor{
		defaultMemoQty: defaultMemoQty,
		memoRepo:       memoRepo,
		tgBot:          tgBot,
	}
}
