package command

import (
	"errors"

	lcerrors "memo/internal/errors"
)

var (
	errNoCmdMessage = errors.New("no message in command")
	errNeedStartCmd = lcerrors.NewReplyError("please, run '/start' command and repeat your request")
	errEmptyPayload = errors.New("empty payload")
	errEmptySender  = errors.New("empty message sender")
)
