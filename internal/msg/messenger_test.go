package msg_test

import (
	"testing"

	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestDummyMessenger(t *testing.T) {
	from := types.PhoneNumber("123")
	to := types.PhoneNumber("457")
	body := "a boring message"
	message := msg.MakeMessage(from, to, body)

	messenger := msg.MakeDummyMessenger()

	r, err := messenger.Send(message)
	assert.Equal(t, "", r)
	assert.NoError(t, err)
}
