package telegram

import (
	"github.com/stretchr/testify/mock"
)

// MockDispatcher is a mocked dispatcher.
type MockDispatcher struct {
	mock.Mock
}

// Send accepts the message to be dispatched and does actually nothing.
func (m *MockDispatcher) Send(message string) error {
	args := m.Called(message)
	return args.Error(0)
}
