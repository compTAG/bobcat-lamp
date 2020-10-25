package msg

import "github.com/comptag/bobcat-lamp/internal/types"

type MessageFactory struct {
	from types.PhoneNumber
}

func MakeMessageFactory(from types.PhoneNumber) MessageFactory {
	return MessageFactory{from}
}

func (f MessageFactory) Create(result types.LabResult) Message {
	body := f.msgForNegative()
	if result.IsPositive() {
		body = f.msgForPositive()
	}

	return MakeMessage(f.from, result.CellPhoneNumber(), body)
}

func (f MessageFactory) msgForPositive() string {
	return "You've got the rona, here is what to do"
}

func (f MessageFactory) msgForNegative() string {
	return "Good news, your test was negative"
}
