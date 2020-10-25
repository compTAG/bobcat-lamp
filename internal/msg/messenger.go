package msg

import (
	"log"

	"github.com/comptag/bobcat-lamp/internal/types"
)

type Messenger interface {
	Send(Message) (string, error)
}

type DummyMessenger struct {
	logger *log.Logger
}

func MakeDummyMessenger(logger *log.Logger) DummyMessenger {
	return DummyMessenger{logger}
}

func (m DummyMessenger) Send(message Message) (string, error) {
	m.logger.Printf("DummyMessenger: (%s) %s", message.To().FullName(), message.Body())
	return "", nil
}
