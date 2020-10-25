package msg_test

import (
	"testing"

	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestMessageGetter(t *testing.T) {
	to := types.Patient{}
	body := "hello world"

	message := msg.MakeMessage(to, body)
	assert.Equal(t, to, message.To())
	assert.Equal(t, body, message.Body())

}
