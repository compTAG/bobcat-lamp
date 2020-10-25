package msg

import "github.com/comptag/bobcat-lamp/internal/types"

type MessageFactory struct {
}

func MakeMessageFactory() MessageFactory {
	return MessageFactory{}
}

func (f MessageFactory) Create(result types.LabResult) Message {
	body := f.msgForNegative()
	if result.IsPositive() {
		body = f.msgForPositive()
	}

	return MakeMessage(result.Patient(), body)
}

func (f MessageFactory) msgForPositive() string {
	return "You've got the rona, here is what to do"
}

func (f MessageFactory) msgForNegative() string {
	return "Good news, your test was negative"
}
