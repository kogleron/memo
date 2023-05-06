package command

import (
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/memo"
	"memo/internal/telegram"
	"memo/internal/user"
)

func NewDeleteExecutor(tgBot telegram.BotAPI, memoRepo memo.Repository, userRepo user.Repository) *DeleteExecutor {
	return &DeleteExecutor{
		tgBot:    tgBot,
		memoRepo: memoRepo,
		userRepo: userRepo,
	}
}

type DeleteExecutor struct {
	tgBot    telegram.BotAPI
	memoRepo memo.Repository
	userRepo user.Repository
}

func (e *DeleteExecutor) Supports(cmd Command) bool {
	return cmd.Name == "delete"
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
		return e.replyTo(cmd.Message, "not found")
	}

	err = e.memoRepo.Delete(memo)
	if err != nil {
		return err
	}

	return e.replyTo(cmd.Message, "deleted")
}

func (e *DeleteExecutor) replyTo(message *tgbotapi.Message, text string) error {
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

func (e *DeleteExecutor) getUser(message *tgbotapi.Message) (*user.User, error) {
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
