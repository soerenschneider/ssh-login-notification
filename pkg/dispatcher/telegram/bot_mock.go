package telegram

import (
	"github.com/stretchr/testify/mock"
)

// MockDispatcher is a mocked dispatcher.
type MockDispatcher struct {
	mock.Mock
}

func (m *MockDispatcher) Send(message string) error {
	args := m.Called(message)
	return args.Error(0)
}
