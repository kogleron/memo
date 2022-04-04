package command

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/user"
)

type StartExecutor struct {
	userRepo *user.Repository
	tgBot    *tgbotapi.BotAPI
}

func (e *StartExecutor) Supports(cmd Command) bool {
	return cmd.Name == "start"
}

func (e *StartExecutor) Run(cmd Command) error {
	if cmd.Message == nil {
		return errNoCmdMessage
	}

	userEntry := e.userRepo.FindByTgAccount(cmd.Message.From.UserName)

	if userEntry == nil {
		userEntry = &user.User{
			TgAccount: cmd.Message.From.UserName,
			TgChatID:  cmd.Message.Chat.ID,
		}
		e.userRepo.Create(userEntry)
	} else {
		userEntry.TgChatID = cmd.Message.Chat.ID
		e.userRepo.Save(userEntry)
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

func NewStartExecutor(userRepo *user.Repository, tgBot *tgbotapi.BotAPI) *StartExecutor {
	return &StartExecutor{
		userRepo: userRepo,
		tgBot:    tgBot,
	}
}
