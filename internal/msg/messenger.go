package msg

import (
	"errors"
	"log"
	"time"

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
	client       SmsClient
	from         types.PhoneNumber
	pollInterval time.Duration
	maxTries     int
}

func MakeSmsMessenger(
	client SmsClient,
	from types.PhoneNumber,
	pollInterval time.Duration,
	maxTries int,
) SmsMessenger {
	return SmsMessenger{client, from, pollInterval, maxTries}
}

func (m SmsMessenger) Send(message Message) (string, error) {
	response, err := m.client.Send(
		m.from,
		message.To().CellPhoneNumber(),
		message.Body(),
	)

	if err != nil {
		return "", err
	}
	sid := response.Sid

	// See https://www.twilio.com/docs/sms/api/message-resource#message-status-values
	// for more on twilio's message status
	for i := 0; i < m.maxTries; i++ {
		response, err = m.client.Get(sid)
		if err == nil {
			switch response.Status {
			case "delivered", "read":
				return sid, nil
			case "failed", "undelivered":
				return sid, errors.New("Message failed to send")
			}
		}
		time.Sleep(m.pollInterval)
	}

	return sid, errors.New("Max number of tries exceeded")
}
