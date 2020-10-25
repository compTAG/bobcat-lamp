package msg

import "github.com/comptag/bobcat-lamp/internal/types"

type Message struct {
	from types.PhoneNumber
	to   types.PhoneNumber
	body string
}

func MakeMessage(from, to types.PhoneNumber, body string) Message {
	return Message{from, to, body}
}

func (m Message) From() types.PhoneNumber {
	return m.from
}

func (m Message) To() types.PhoneNumber {
	return m.to
}

func (m Message) Body() string {
	return m.body
}
