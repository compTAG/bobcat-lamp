package msg

import "github.com/comptag/bobcat-lamp/internal/types"

type Message struct {
	to   types.PhoneNumber
	body string
}

func MakeMessage(to types.PhoneNumber, body string) Message {
	return Message{to, body}
}

func (m Message) To() types.PhoneNumber {
	return m.to
}

func (m Message) Body() string {
	return m.body
}
