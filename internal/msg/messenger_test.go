package msg_test

import (
	"errors"
	"log"
	"os"
	"testing"
	"time"

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

func makeResponse(sid, status string) *msg.SmsResponse {
	return &(msg.SmsResponse{Sid: sid, Status: status})
}

func TestSmsMessenger(t *testing.T) {

	cell := types.MakePhoneNumber("1112223333")
	to := types.MakePatient("", "", cell)
	body := "a boring message"
	message := msg.MakeMessage(to, body)

	serverPhone := types.MakePhoneNumber("4445556666")

	maxTries := 5
	pollInterval := time.Millisecond

	sid := "a-sid"

	t.Run("fail to submit sms", func(t *testing.T) {
		client := new(mocks.SmsClient)
		err := errors.New("an error")
		client.On("Send", serverPhone, cell, body).Return(nil, err)

		messenger := msg.MakeSmsMessenger(client, serverPhone, pollInterval, maxTries)

		r, err := messenger.Send(message)
		client.AssertExpectations(t)
		assert.Equal(t, "", r)
		assert.Error(t, err)
	})

	t.Run("send success on first try", func(t *testing.T) {
		client := new(mocks.SmsClient)
		client.
			On("Send", serverPhone, cell, body).
			Return(makeResponse(sid, "queued"), nil)
		client.
			On("Get", sid).
			Return(makeResponse(sid, "delivered"), nil)

		messenger := msg.MakeSmsMessenger(client, serverPhone, pollInterval, maxTries)

		r, err := messenger.Send(message)
		client.AssertExpectations(t)
		assert.Equal(t, sid, r)
		assert.NoError(t, err)
	})

	t.Run("send success as read", func(t *testing.T) {
		client := new(mocks.SmsClient)
		client.
			On("Send", serverPhone, cell, body).
			Return(makeResponse(sid, "queued"), nil)
		client.
			On("Get", sid).
			Return(makeResponse(sid, "read"), nil)

		messenger := msg.MakeSmsMessenger(client, serverPhone, pollInterval, maxTries)

		r, err := messenger.Send(message)
		client.AssertExpectations(t)
		assert.Equal(t, sid, r)
		assert.NoError(t, err)
	})

	t.Run("send success after waiting a bit", func(t *testing.T) {
		client := new(mocks.SmsClient)
		client.
			On("Send", serverPhone, cell, body).
			Return(makeResponse(sid, "accepted"), nil)
		client.
			On("Get", sid).
			Return(makeResponse(sid, "queued"), nil).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "sending"), nil).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "sent"), nil).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "delivered"), nil).
			Once()

		messenger := msg.MakeSmsMessenger(client, serverPhone, pollInterval, maxTries)

		r, err := messenger.Send(message)
		client.AssertExpectations(t)
		assert.Equal(t, sid, r)
		assert.NoError(t, err)
	})

	t.Run("get request intermittentently fails", func(t *testing.T) {
		client := new(mocks.SmsClient)
		client.
			On("Send", serverPhone, cell, body).
			Return(makeResponse(sid, "queued"), nil)
		client.
			On("Get", sid).
			Return(makeResponse(sid, "sending"), nil).
			Once()
		client.
			On("Get", sid).
			Return(nil, errors.New("error")).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "sent"), nil).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "delivered"), nil).
			Once()

		messenger := msg.MakeSmsMessenger(client, serverPhone, pollInterval, maxTries)

		r, err := messenger.Send(message)
		client.AssertExpectations(t)
		assert.Equal(t, sid, r)
		assert.NoError(t, err)
	})

	t.Run("message fails", func(t *testing.T) {
		client := new(mocks.SmsClient)
		client.
			On("Send", serverPhone, cell, body).
			Return(makeResponse(sid, "accepted"), nil)
		client.
			On("Get", sid).
			Return(makeResponse(sid, "queued"), nil).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "failed"), nil).
			Once()

		messenger := msg.MakeSmsMessenger(client, serverPhone, pollInterval, maxTries)

		r, err := messenger.Send(message)
		client.AssertExpectations(t)
		assert.Equal(t, sid, r)
		assert.Error(t, err)

	})

	t.Run("message becomes undelevered", func(t *testing.T) {
		client := new(mocks.SmsClient)
		client.
			On("Send", serverPhone, cell, body).
			Return(makeResponse(sid, "accepted"), nil)
		client.
			On("Get", sid).
			Return(makeResponse(sid, "queued"), nil).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "undelivered"), nil).
			Once()

		messenger := msg.MakeSmsMessenger(client, serverPhone, pollInterval, maxTries)

		r, err := messenger.Send(message)
		client.AssertExpectations(t)
		assert.Equal(t, sid, r)
		assert.Error(t, err)

	})

	t.Run("send times out", func(t *testing.T) {
		client := new(mocks.SmsClient)
		client.
			On("Send", serverPhone, cell, body).
			Return(makeResponse(sid, "accepted"), nil)
		client.
			On("Get", sid).
			Return(makeResponse(sid, "queued"), nil).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "queued"), nil).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "sending"), nil).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "sent"), nil).
			Once()
		client.
			On("Get", sid).
			Return(makeResponse(sid, "sent"), nil).
			Once()

		messenger := msg.MakeSmsMessenger(client, serverPhone, pollInterval, maxTries)

		r, err := messenger.Send(message)
		client.AssertExpectations(t)
		assert.Equal(t, sid, r)
		assert.Error(t, err)
	})
}
