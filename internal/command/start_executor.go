package command

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/telegram"
	"memo/internal/user"
)

type StartExecutor struct {
	userRepo user.Repository
	tgBot    telegram.BotAPI
}

func (e StartExecutor) GetName() string {
	return "start"
}

func (e StartExecutor) GetDescription() string {
	return "starts work with a user"
}

func (e StartExecutor) Supports(cmd Command) bool {
	return cmd.Name == e.GetName()
}

func (e *StartExecutor) Run(cmd Command) error {
	if cmd.Message == nil {
		return errNoCmdMessage
	}

	userEntry, err := e.userRepo.FindByTgAccount(cmd.Message.From.UserName)
	if err != nil {
		return err
	}

	if userEntry == nil {
		userEntry = &user.User{
			TgAccount: cmd.Message.From.UserName,
			TgChatID:  cmd.Message.Chat.ID,
		}

		err := e.userRepo.Create(userEntry)
		if err != nil {
			return err
		}
	} else {
		userEntry.TgChatID = cmd.Message.Chat.ID
		err := e.userRepo.Save(userEntry)
		if err != nil {
			return err
		}
	}

	msg := tgbotapi.NewMessage(
		cmd.Message.Chat.ID,
		"done",
	)
	msg.ReplyToMessageID = cmd.Message.MessageID

	_, err = e.tgBot.Send(msg)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func NewStartExecutor(userRepo user.Repository, tgBot telegram.BotAPI) *StartExecutor {
	return &StartExecutor{
		userRepo: userRepo,
		tgBot:    tgBot,
	}
}
