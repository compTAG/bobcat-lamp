package lab_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/comptag/bobcat-lamp/internal/lab"
	"github.com/comptag/bobcat-lamp/internal/types"
)

func TestResultGetter(t *testing.T) {
	phone := types.MakePhoneNumber("123")
	result := lab.MakeResult("abc", "bob c lamptest", phone, false)

	assert.Equal(t, result.Id(), "abc")
	assert.Equal(t, result.FullName(), "bob c lamptest")
	assert.Equal(t, result.CellPhoneNumber(), phone)
	assert.Equal(t, result.IsPositive(), false)
}
