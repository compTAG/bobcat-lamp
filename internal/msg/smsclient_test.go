package msg_test

import (
	"testing"

	"github.com/comptag/bobcat-lamp/internal/msg"
)

func TestSmsClientInterface(t *testing.T) {
	// if result is not a lab result this code will cause compilation to fail
	var _ msg.SmsClient = msg.MakeTwilioClient("", "")
}
