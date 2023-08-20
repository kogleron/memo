package telegram

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/configs"
	"memo/internal/api/telegram/command"
	lcerrors "memo/internal/pkg/errors"
)

func NewPollingBot(
	tgBot *tgbotapi.BotAPI,
	tgConfig configs.TelegramConfig,
	cmdParser command.Parser,
	cmdExecutors []command.Executor,
	shutdownAfterTimeout bool,
) PollingBot {
	return &pollingBot{
		tgBot:                tgBot,
		tgConfig:             tgConfig,
		cmdParser:            cmdParser,
		cmdExecutors:         cmdExecutors,
		shutdownAfterTimeout: shutdownAfterTimeout,
	}
}

type PollingBot interface {
	Run()
}

type pollingBot struct {
	tgBot                *tgbotapi.BotAPI
	tgConfig             configs.TelegramConfig
	cmdParser            command.Parser
	cmdExecutors         []command.Executor
	shutdownAfterTimeout bool
}

func (b *pollingBot) Run() {
	updateConf := tgbotapi.NewUpdate(0)
	updateConf.Timeout = 3
	updateConf.AllowedUpdates = append(updateConf.AllowedUpdates, "message")
	updatesCh := b.tgBot.GetUpdatesChan(updateConf)

	if b.shutdownAfterTimeout {
		go func() {
			time.Sleep(time.Second * time.Duration(updateConf.Timeout+1))
			b.tgBot.StopReceivingUpdates()
			log.Println("stopping...")
		}()
	}

	for update := range updatesCh {
		if !b.isAllowedUpdate(&update) {
			log.Printf("message from a not authorized account %s\n", update.Message.From.UserName)
			b.replyError(lcerrors.NewReplyError("you are not authorized"), update.Message)

			continue
		}

		err := b.processUpdate(&update)
		if err != nil {
			b.replyError(err, update.Message)
		}
	}

	log.Println("shutdowning...")
}

func (b *pollingBot) isAllowedUpdate(update *tgbotapi.Update) bool {
	_, allowed := b.tgConfig.AllowedAccountsMap[update.Message.From.UserName]

	return allowed
}

func (b *pollingBot) processUpdate(update *tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	cmd, err := b.parseCommand(update.Message)
	if err != nil {
		return err
	}

	for _, cmdExecutor := range b.cmdExecutors {
		if !cmdExecutor.Supports(*cmd) {
			continue
		}

		err = cmdExecutor.Run(*cmd)
		if err != nil {
			return err
		}

		return nil
	}

	return lcerrors.NewReplyError("not supported command '" + cmd.Name + "'")
}

func (b pollingBot) replyError(err error, message *tgbotapi.Message) {
	log.Println(err)

	_, needReply := err.(*lcerrors.ReplyError) //nolint: errorlint
	if !needReply {
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		err.Error(),
	)
	msg.ReplyToMessageID = message.MessageID

	_, err = b.tgBot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b pollingBot) parseCommand(message *tgbotapi.Message) (*command.Command, error) {
	if !b.cmdParser.IsCommand(message) {
		return nil, lcerrors.NewReplyError("not a command '" + message.Text + "'")
	}

	cmd, err := b.cmdParser.ParseCommand(message)
	if err != nil {
		log.Print(err)

		return nil, lcerrors.NewReplyError("failed to parse command")
	}

	return cmd, nil
}
