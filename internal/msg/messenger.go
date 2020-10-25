package msg

type Messenger interface {
	Send(Message) (string, error)
}

type DummyMessenger struct {
}

func MakeDummyMessenger() DummyMessenger {
	return DummyMessenger{}
}

func (m *DummyMessenger) Send(message Message) (string, error) {
	return "", nil
}
