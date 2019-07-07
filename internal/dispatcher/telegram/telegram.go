package telegram

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"sshnot/internal"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type telegramBot struct {
	bot  		*tgbotapi.BotAPI
	receiver	int64
}

func New(option *internal.Options) (*telegramBot, error) {
	if nil == option {
		return nil, errors.New("No configuration supplied.")
	}

	client := telegramBot{receiver: option.TelegramId}
	var err error

	client.bot, err = tgbotapi.NewBotAPI(option.TelegramToken)
	if err != nil {
		return nil, err
	}

	log.Printf("Authorized on account %s", client.bot.Self.UserName)
	return &client, nil
}

func (b *telegramBot) Send(message string) {
	msg := tgbotapi.NewMessage(b.receiver, message)
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Errorf("Can't send message: %v", err)
	}
}