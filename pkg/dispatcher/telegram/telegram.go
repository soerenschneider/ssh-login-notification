package telegram

import (
	"errors"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"sshnot/internal"
)

type telegramBot struct {
	bot      *tgbotapi.BotAPI
	receiver int64
}

// NewTelegramBot instantiates a new telegram bot.
func NewTelegramBot(option *internal.Options) (*telegramBot, error) {
	if nil == option {
		return nil, errors.New("No configuration supplied.")
	}

	client := telegramBot{receiver: option.TelegramId}
	var err error

	client.bot, err = tgbotapi.NewBotAPI(option.TelegramToken)
	if err != nil {
		return nil, err
	}

	log.Debugf("Authorized on account %s", client.bot.Self.UserName)
	return &client, nil
}

// Send accepts the message to be send to the user and dispatches it via telegram.
func (b *telegramBot) Send(message string) {
	msg := tgbotapi.NewMessage(b.receiver, message)
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Errorf("Can't send message: %v", err)
	}
}
