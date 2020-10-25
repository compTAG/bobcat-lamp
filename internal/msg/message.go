package msg

import "github.com/comptag/bobcat-lamp/internal/types"

type Message struct {
	to   types.Patient
	body string
}

func MakeMessage(to types.Patient, body string) Message {
	return Message{to, body}
}

func (m Message) To() types.Patient {
	return m.to
}

func (m Message) Body() string {
	return m.body
}
