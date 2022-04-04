package apps

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"memo/internal/command"
	"memo/internal/configs"
)

type PollingBot struct {
	tgBot        *tgbotapi.BotAPI
	tgConfig     configs.TelegramConfig
	cmdParser    *command.Parser
	cmdExecutors []command.Executor
}

func (b *PollingBot) Run() {
	updateConf := tgbotapi.NewUpdate(0)
	updateConf.Timeout = 60
	updateConf.AllowedUpdates = append(updateConf.AllowedUpdates, "message")
	updatesCh := b.tgBot.GetUpdatesChan(updateConf)

	for update := range updatesCh {
		if !b.isAllowedUpdate(&update) {
			continue
		}
		b.processUpdate(&update)
	}
}

func (b *PollingBot) isAllowedUpdate(update *tgbotapi.Update) bool {
	_, allowed := b.tgConfig.AllowedAccountsMap[update.Message.From.UserName]

	if !allowed {
		log.Printf("Message from a not authorized account %s\n", update.Message.From.UserName)
	}

	return allowed
}

func (b *PollingBot) processUpdate(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if !b.cmdParser.IsCommand(update.Message) {
		log.Printf("not a command: %s\n", update.Message.Text)
		return
	}

	cmd, err := b.cmdParser.ParseCommand(update.Message.Text)
	if err != nil {
		log.Println(err)
		return
	}
	if cmd == nil {
		return
	}

	cmdWasExecuted := false
	for _, cmdExecutor := range b.cmdExecutors {
		if !cmdExecutor.Supports(*cmd) {
			continue
		}
		err = cmdExecutor.Run(*cmd)
		if err != nil {
			log.Println(err)
		}
		cmdWasExecuted = true
		break
	}

	if cmdWasExecuted {
		b.onSuccessExecution(update)
		return
	}
	b.onFailedExecution(update, *cmd)
}

func (b *PollingBot) onSuccessExecution(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"done",
	)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := b.tgBot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b *PollingBot) onFailedExecution(update *tgbotapi.Update, cmd command.Command) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I do not have executor for the given command: "+cmd.Name)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := b.tgBot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func NewPollingBot(
	tgBot *tgbotapi.BotAPI,
	tgConfig configs.TelegramConfig,
	cmdParser *command.Parser,
	cmdExecutors []command.Executor,
) *PollingBot {
	return &PollingBot{
		tgBot:        tgBot,
		tgConfig:     tgConfig,
		cmdParser:    cmdParser,
		cmdExecutors: cmdExecutors,
	}
}