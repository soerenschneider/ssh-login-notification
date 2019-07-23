package telegram

import (
	"github.com/stretchr/testify/mock"
	"sshnot/internal"
)

type telegramMock struct {
	mock.Mock
	receiver int64
}

func NewTelegramMock(option *internal.Options) (*telegramMock, error) {
	return &telegramMock{receiver: option.TelegramId}, nil
}

func (b *telegramMock) Send(message string) error {
	args := b.Called(message)
	return args.Error(0)
}
