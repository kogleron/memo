package command

import "errors"

var (
	errEmptyMessage = errors.New("empty message")
	errNoCmdMessafe = errors.New("no message in command")
)
