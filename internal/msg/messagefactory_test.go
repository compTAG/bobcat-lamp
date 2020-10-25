package msg_test

import (
	"testing"

	"github.com/comptag/bobcat-lamp/internal/lab"
	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestMessageFactoryCreate(t *testing.T) {
	from := types.MakePhoneNumber("1115557777")

	cases := []struct {
		name         string
		to           types.PhoneNumber
		result       bool
		expectedBody string
	}{
		{
			"test is negative",
			types.MakePhoneNumber("5555555555"),
			false,
			"Good news, your test was negative",
		}, {
			"test is positive",
			types.MakePhoneNumber("6666666666"),
			true,
			"You've got the rona, here is what to do",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := lab.MakeResult("id", "Name X", tc.to, tc.result)

			factory := msg.MakeMessageFactory(from)
			message := factory.Create(result)

			assert.Equal(t, from, message.From())
			assert.Equal(t, tc.to, message.To())
			assert.Equal(t, tc.expectedBody, message.Body())
		})
	}
}
