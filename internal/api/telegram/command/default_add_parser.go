package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/domain"
)

func NewDefaultAddParser(
	next Parser,
	userRepo domain.UserRepository,
) Parser {
	return &DefaultAddParser{
		next:     next,
		userRepo: userRepo,
	}
}

type DefaultAddParser struct {
	next     Parser
	userRepo domain.UserRepository
}

func (p DefaultAddParser) IsCommand(message *tgbotapi.Message) bool {
	return p.next.IsCommand(message)
}

func (p DefaultAddParser) ParseCommand(message *tgbotapi.Message) (*Command, error) {
	command, err := p.next.ParseCommand(message)
	if err != nil {
		return nil, err
	}

	if command != nil {
		return command, nil
	}

	user, err := p.getUser(message)
	if err != nil {
		return nil, err
	}

	return &Command{
		Name:    (new(AddExecutor)).GetName(),
		Payload: message.Text,
		Message: message,
		Sender:  user,
	}, nil
}

func (p DefaultAddParser) getUser(message *tgbotapi.Message) (user *domain.User, err error) {
	if message == nil || message.From == nil {
		return
	}

	return p.userRepo.FindByTgAccount(message.From.UserName)
}
