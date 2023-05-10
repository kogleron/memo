package telegram

import (
	"errors"
)

var (
	ErrEmptyMessage = errors.New("empty message")
	ErrNoChatID     = errors.New("no chat id")
)
