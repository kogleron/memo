package apps

import (
	"errors"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/domain"
	"memo/internal/pkg/telegram"
)

type RandCommand struct {
	defaultMemoQty uint
	memoRepo       domain.MemoRepository
	userRepo       domain.UserRepository
	tgBot          telegram.BotAPI
}

func (c *RandCommand) Run() {
	users, err := c.userRepo.FindAll()
	if err != nil {
		log.Println(err)

		return
	}

	for i := range users {
		user := &users[i]

		memos, err := c.memoRepo.Rand(c.defaultMemoQty, user)
		if err != nil {
			log.Println(err)

			continue
		}

		if len(memos) == 0 {
			continue
		}

		err = c.sendMemos(user, memos)
		if err != nil {
			log.Println(err)
		}
	}
}

func (c *RandCommand) sendMemos(user *domain.User, memos []domain.Memo) error {
	if user.TgChatID == 0 {
		return errors.New("no chat id for user " + user.TgAccount) //nolint: goerr113
	}

	var err error

	for i := range memos {
		msg := tgbotapi.NewMessage(
			user.TgChatID,
			fmt.Sprintf("#<b>%d</b>\n%s", memos[i].ID, memos[i].Text),
		)
		msg.DisableNotification = true
		msg.ParseMode = "html"

		_, err := c.tgBot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}

	return err
}

func NewRandCommand(
	defaultMemoQty uint,
	memoRepo domain.MemoRepository,
	userRepo domain.UserRepository,
	tgBot telegram.BotAPI,
) *RandCommand {
	return &RandCommand{
		defaultMemoQty: defaultMemoQty,
		memoRepo:       memoRepo,
		userRepo:       userRepo,
		tgBot:          tgBot,
	}
}
