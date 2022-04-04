package apps

import (
	"errors"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/memo"
	"memo/internal/user"
)

type RandCommand struct {
	defaultMemoQty uint
	memoRepo       *memo.Repository
	userRepo       *user.Repository
	tgBot          *tgbotapi.BotAPI
}

func (c *RandCommand) Run() {
	users := c.userRepo.FindAll()

	for _, user := range users {
		memos := c.memoRepo.Rand(c.defaultMemoQty)
		if len(memos) == 0 {
			continue
		}

		err := c.sendMemos(user, memos)
		if err != nil {
			log.Println(err)
		}
	}
}

func (c *RandCommand) sendMemos(user user.User, memos []memo.Memo) error {
	if user.TgChatID == 0 {
		return errors.New("no chat id for user " + user.TgAccount) //nolint: goerr113
	}

	replyText := ""

	for i := len(memos) - 1; i >= 0; i-- {
		replyText += memos[i].Text
		if i > 0 {
			replyText += "\n\n"
		}
	}

	msg := tgbotapi.NewMessage(
		user.TgChatID,
		replyText,
	)

	_, err := c.tgBot.Send(msg)

	return err
}

func NewRandCommand(
	defaultMemoQty uint,
	memoRepo *memo.Repository,
	userRepo *user.Repository,
	tgBot *tgbotapi.BotAPI,
) *RandCommand {
	return &RandCommand{
		defaultMemoQty: defaultMemoQty,
		memoRepo:       memoRepo,
		userRepo:       userRepo,
		tgBot:          tgBot,
	}
}
