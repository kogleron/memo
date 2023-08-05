package command

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/domain"
	"memo/internal/pkg/telegram"
)

type RandExecutor struct {
	defaultMemoQty uint
	memoRepo       domain.MemoRepository
	tgBot          telegram.BotAPI
}

func (e RandExecutor) GetName() string {
	return "rand"
}

func (e RandExecutor) GetDescription() string {
	return "retrieves " + strconv.FormatUint(uint64(e.defaultMemoQty), 10) + " random memos"
}

func (e RandExecutor) Supports(cmd Command) bool {
	return cmd.Name == e.GetName()
}

func (e *RandExecutor) Run(cmd Command) error {
	if cmd.Message == nil {
		return errNoCmdMessage
	}

	if cmd.Sender == nil {
		return errNeedStartCmd
	}

	message := cmd.Message

	memos, err := e.memoRepo.Rand(e.defaultMemoQty, cmd.Sender)
	if err != nil {
		return err
	}

	if len(memos) == 0 {
		return e.onNoMemos(message)
	}

	return e.sendMemos(message, memos)
}

func (e *RandExecutor) sendMemos(message *tgbotapi.Message, memos []domain.Memo) error {
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

func NewRandExecutor(
	defaultMemoQty uint,
	memoRepo domain.MemoRepository,
	tgBot telegram.BotAPI,
) *RandExecutor {
	return &RandExecutor{
		defaultMemoQty: defaultMemoQty,
		memoRepo:       memoRepo,
		tgBot:          tgBot,
	}
}
