package telegram

import (
	"errors"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/prometheus/common/log"
	"sshnot/internal"
)

type telegramBot struct {
	bot      *tgbotapi.BotAPI
	receiver int64
	token    string
}

// NewTelegramBot instantiates a new telegram bot.
func NewTelegramBot(option *internal.Options) (*telegramBot, error) {
	if nil == option {
		return nil, errors.New("No configuration supplied.")
	}

	client := telegramBot{receiver: option.TelegramId, token: option.TelegramToken}
	return &client, nil
}

// Send accepts the message to be send to the user and dispatches it via telegram.
func (b *telegramBot) Send(message string) error {
	var err error

	b.bot, err = tgbotapi.NewBotAPI(b.token)
	if err != nil {
		return err
	}

	log.Debugf("Authorized on account %s", b.bot.Self.UserName)
	msg := tgbotapi.NewMessage(b.receiver, message)
	_, err = b.bot.Send(msg)
	return err
}
