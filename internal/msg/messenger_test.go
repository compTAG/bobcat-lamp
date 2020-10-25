package msg_test

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/internal/types"
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
