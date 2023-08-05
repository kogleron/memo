package command

import (
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/domain"
	"memo/internal/pkg/telegram"
)

func NewParser(
	userRepo domain.UserRepository,
) Parser {
	return parser{
		userRepo: userRepo,
	}
}

type Parser interface {
	IsCommand(message *tgbotapi.Message) bool
	ParseCommand(message *tgbotapi.Message) (*Command, error)
}

type parser struct {
	userRepo domain.UserRepository
}

func (p parser) IsCommand(message *tgbotapi.Message) bool {
	return message != nil
}

func (p parser) ParseCommand(message *tgbotapi.Message) (*Command, error) {
	if message == nil {
		return nil, telegram.ErrEmptyMessage
	}

	text := strings.Trim(message.Text, " ")
	if len(text) == 0 {
		return nil, telegram.ErrEmptyMessage
	}

	reg := regexp.MustCompile(`^/([a-z]+)\s*(.*?)\s*$`)

	matches := reg.FindAllStringSubmatch(text, -1)
	if len(matches) == 0 {
		return nil, nil //nolint
	}

	user, err := p.getUser(message)
	if err != nil {
		return nil, err
	}

	return &Command{
		Name:    matches[0][1],
		Payload: matches[0][2],
		Message: message,
		Sender:  user,
	}, nil
}

func (p parser) getUser(message *tgbotapi.Message) (user *domain.User, err error) {
	if message == nil || message.From == nil {
		return
	}

	return p.userRepo.FindByTgAccount(message.From.UserName)
}
