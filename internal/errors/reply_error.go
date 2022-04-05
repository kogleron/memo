package errors

import "errors"

type ReplyError struct {
	error error
}

func (e *ReplyError) Error() string {
	return e.error.Error()
}

func NewReplyError(text string) *ReplyError {
	return &ReplyError{
		error: errors.New(text), //nolint
	}
}
