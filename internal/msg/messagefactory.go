package msg

import "github.com/comptag/bobcat-lamp/internal/types"

type MessageFactory interface {
	Create(result types.LabResult) Message
}

type StaticMessageFactory struct {
}

func MakeStaticMessageFactory() StaticMessageFactory {
	return StaticMessageFactory{}
}

func (f StaticMessageFactory) Create(result types.LabResult) Message {
	body := f.msgForNegative()
	if result.IsPositive() {
		body = f.msgForPositive()
	}

	return MakeMessage(result.Patient(), body)
}

func (f StaticMessageFactory) msgForPositive() string {
	return "You've got the rona, here is what to do"
}

func (f StaticMessageFactory) msgForNegative() string {
	return "Good news, your test was negative"
}
