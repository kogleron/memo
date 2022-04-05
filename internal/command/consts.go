package command

import (
	"errors"

	lcerrors "memo/internal/errors"
)

var (
	errEmptyMessage = errors.New("empty message")
	errNoCmdMessage = errors.New("no message in command")
	errNeedStartCmd = lcerrors.NewReplyError("please, run '/start' command and repeat your request")
)
