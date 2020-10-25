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

type SmsMessenger struct {
	client SmsClient
	from   types.PhoneNumber
}

func MakeSmsMessenger(client SmsClient, from types.PhoneNumber) SmsMessenger {
	return SmsMessenger{client, from}
}

func (m SmsMessenger) Send(message Message) (string, error) {
	response, err := m.client.Send(
		m.from,
		message.To().CellPhoneNumber(),
		message.Body(),
	)

	sid := ""
	if err == nil {
		sid = response.Sid
	}
	return sid, err
}
