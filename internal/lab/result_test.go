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

	assert.Equal(t, "abc", result.Id())
	assert.Equal(t, "bob c lamptest", result.FullName())
	assert.Equal(t, phone, result.CellPhoneNumber())
	assert.False(t, result.IsPositive())
}
