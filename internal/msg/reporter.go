package msg

import "github.com/comptag/bobcat-lamp/internal/types"

type Reporter struct {
	messenger  Messenger
	msgFactory MessageFactory
}

func MakeReporterWithMessengerAndMsgFactory(
	messenger Messenger,
	factory MessageFactory,
) Reporter {
	return Reporter{messenger, factory}
}

func MakeReporterWithMessenger(messenger Messenger) Reporter {
	msgFactory := MakeStaticMessageFactory()
	return Reporter{messenger, msgFactory}
}

func MakeDummyReporter() Reporter {
	dummyBackend := MakeDummyMessenger()
	return MakeReporterWithMessenger(dummyBackend)
}

func (r *Reporter) Report(result types.LabResult) (string, error) {
	message := r.msgFactory.Create(result)
	return r.messenger.Send(message)
}
