package msg_test

import (
	"testing"

	"github.com/comptag/bobcat-lamp/internal/lab"
	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/mocks"
)

func TestReporterReport(t *testing.T) {
	labResult := lab.Result{}

	message := msg.Message{}
	msgFactory := new(mocks.MessageFactory)
	msgFactory.On("Create", labResult).Return(message)

	messenger := new(mocks.Messenger)
	messenger.On("Send", message).Return("", nil)

	reporter := msg.MakeReporterWithMessengerAndMsgFactory(messenger, msgFactory)
	reporter.Report(labResult)

	msgFactory.AssertExpectations(t)
	messenger.AssertExpectations(t)
}
