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
			log.Println("Stopping...")
		}()
	}

	for update := range updatesCh {
		if !b.isAllowedUpdate(&update) {
			continue
		}

		b.processUpdate(&update)
	}

	log.Println("Shutdowning...")
}

func (b *pollingBot) isAllowedUpdate(update *tgbotapi.Update) bool {
	_, allowed := b.tgConfig.AllowedAccountsMap[update.Message.From.UserName]

	if !allowed {
		log.Printf("Message from a not authorized account %s\n", update.Message.From.UserName)
	}

	return allowed
}

func (b *pollingBot) processUpdate(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if !b.cmdParser.IsCommand(update.Message) {
		log.Printf("Not a command: %s\n", update.Message.Text)

		return
	}

	cmd, err := b.cmdParser.ParseCommand(update.Message)
	if err != nil {
		log.Println(err)

		return
	}

	for _, cmdExecutor := range b.cmdExecutors {
		if !cmdExecutor.Supports(*cmd) {
			continue
		}

		err = cmdExecutor.Run(*cmd)
		if err != nil {
			b.onExecutionError(err, update.Message)
		}

		return
	}

	b.onFailedExecution(update, *cmd)
}

func (b *pollingBot) onExecutionError(err error, message *tgbotapi.Message) {
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

func (b *pollingBot) onFailedExecution(update *tgbotapi.Update, cmd command.Command) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Not supported command: "+cmd.Name)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := b.tgBot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
