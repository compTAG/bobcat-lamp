package msg_test

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/internal/types"
	"github.com/comptag/bobcat-lamp/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDummyMessenger(t *testing.T) {
	to := types.Patient{}
	body := "a boring message"
	message := msg.MakeMessage(to, body)

	logger := log.New(os.Stdout, "", log.LstdFlags)
	messenger := msg.MakeDummyMessenger(logger)

	r, err := messenger.Send(message)
	assert.Equal(t, "", r)
	assert.NoError(t, err)
}

func TestSmsMessenger(t *testing.T) {

	cell := types.MakePhoneNumber("1112223333")
	to := types.MakePatient("", "", cell)
	body := "a boring message"
	message := msg.MakeMessage(to, body)

	serverPhone := types.MakePhoneNumber("4445556666")

	t.Run("send success", func(t *testing.T) {
		sid := "a-sid"

		client := new(mocks.SmsClient)
		response := msg.SmsResponse{Sid: sid}
		client.On("Send", serverPhone, cell, body).Return(&response, nil)

		messenger := msg.MakeSmsMessenger(client, serverPhone)

		r, err := messenger.Send(message)
		client.AssertExpectations(t)
		assert.Equal(t, sid, r)
		assert.NoError(t, err)
	})

	t.Run("send fails", func(t *testing.T) {

		client := new(mocks.SmsClient)
		err := errors.New("an error")
		client.On("Send", serverPhone, cell, body).Return(nil, err)

		messenger := msg.MakeSmsMessenger(client, serverPhone)

		r, err := messenger.Send(message)
		client.AssertExpectations(t)
		assert.Equal(t, "", r)
		assert.Error(t, err)
	})
}
