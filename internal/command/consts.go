package command

import "errors"

var (
	errEmptyMessage = errors.New("empty message")
	errNoCmdMessage = errors.New("no message in command")
)
