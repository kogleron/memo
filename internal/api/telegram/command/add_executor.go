package command

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/domain"
	"memo/internal/pkg/telegram"
)

func NewAddExecutor(
	memoRepo domain.MemoRepository,
	tgBot telegram.BotAPI,
) *AddExecutor {
	return &AddExecutor{
		memoRepo: memoRepo,
		tgBot:    tgBot,
	}
}

type AddExecutor struct {
	memoRepo domain.MemoRepository
	tgBot    telegram.BotAPI
}

func (e AddExecutor) GetDescription() string {
	return "creates the memo with a given text. \nFormat: /" + e.GetName() + " [MEMO_TEXT]"
}

func (AddExecutor) GetName() string {
	return "add"
}

func (e AddExecutor) Supports(cmd Command) bool {
	return cmd.Name == e.GetName()
}

func (e *AddExecutor) Run(cmd Command) error {
	if cmd.Message == nil {
		return errNoCmdMessage
	}

	if cmd.Sender == nil {
		return errNeedStartCmd
	}

	memo := &domain.Memo{
		Text:   strings.TrimSpace(cmd.Payload),
		UserID: cmd.Sender.ID,
	}
	if memo.Text == "" {
		return errNoCmdMessage
	}

	err := e.memoRepo.Create(memo)
	if err != nil {
		return err
	}

	if cmd.Message == nil {
		return nil
	}

	msg := tgbotapi.NewMessage(
		cmd.Message.Chat.ID,
		fmt.Sprintf("added with id <b>%d</b>", memo.ID),
	)
	msg.ParseMode = "html"
	msg.ReplyToMessageID = cmd.Message.MessageID
	msg.DisableWebPagePreview = true
	msg.DisableNotification = true

	_, err = e.tgBot.Send(msg)
	if err != nil {
		log.Println(err)
	}

	return nil
}
