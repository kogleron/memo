package command

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

type AddExecutor struct {
	memoRepo memo.Repository
	tgBot    telegram.BotAPI
	userRepo user.Repository
}

func (e *AddExecutor) Supports(cmd Command) bool {
	return cmd.Name == "add"
}

func (e *AddExecutor) Run(cmd Command) error {
	if cmd.Message == nil {
		return errNoCmdMessage
	}

	user, err := e.getUser(cmd.Message.From.UserName)
	if err != nil {
		return err
	}

	memo := &memo.Memo{
		Text:   strings.TrimSpace(cmd.Payload),
		UserID: user.ID,
	}
	if memo.Text == "" {
		return errNoCmdMessage
	}

	err = e.memoRepo.Create(memo)
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

func (e *AddExecutor) getUser(tgAccountName string) (*user.User, error) {
	user, err := e.userRepo.FindByTgAccount(tgAccountName)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errNeedStartCmd
	}

	return user, nil
}

func NewAddExecutor(memoRepo memo.Repository, tgBot telegram.BotAPI, userRepo user.Repository) *AddExecutor {
	return &AddExecutor{
		memoRepo: memoRepo,
		tgBot:    tgBot,
		userRepo: userRepo,
	}
}
