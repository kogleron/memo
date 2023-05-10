package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewReplier(
	tgBot BotAPI,
) Replier {
	return &replier{
		tgBot: tgBot,
	}
}

type Replier interface {
	ReplyTo(message *tgbotapi.Message, text string) error
}

type replier struct {
	tgBot BotAPI
}

func (r *replier) ReplyTo(message *tgbotapi.Message, text string) error {
	if message == nil {
		return ErrEmptyMessage
	}

	if message.Chat == nil {
		return ErrNoChatID
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		text,
	)

	msg.ReplyToMessageID = message.MessageID
	msg.DisableWebPagePreview = true
	msg.ParseMode = tgbotapi.ModeHTML

	_, err := r.tgBot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
