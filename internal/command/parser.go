package command

import (
	"errors"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Parser struct{}

func (p *Parser) IsCommand(message *tgbotapi.Message) bool {
	return message != nil
}

func (p *Parser) ParseCommand(message string) (*Command, error) {
	message = strings.Trim(message, " ")
	if len(message) == 0 {
		return nil, errors.New("empty message")
	}

	return &Command{
		Name:    "add",
		Payload: message,
	}, nil
}

func NewParser() *Parser {
	return &Parser{}
}
