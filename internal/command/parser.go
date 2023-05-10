package command

import (
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/telegram"
)

type Parser struct{}

func (p *Parser) IsCommand(message *tgbotapi.Message) bool {
	return message != nil
}

func (p *Parser) ParseCommand(message *tgbotapi.Message) (*Command, error) {
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
		return &Command{
			Name:    "add",
			Payload: text,
			Message: message,
		}, nil
	}

	return &Command{
		Name:    matches[0][1],
		Payload: matches[0][2],
		Message: message,
	}, nil
}

func NewParser() *Parser {
	return &Parser{}
}
