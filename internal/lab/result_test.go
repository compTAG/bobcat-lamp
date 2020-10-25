package lab_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/comptag/bobcat-lamp/internal/lab"
	"github.com/comptag/bobcat-lamp/internal/types"
)

func TestResultInterface(t *testing.T) {
	// if result is not a lab result this code will cause compilation to fail
	phone := types.MakePhoneNumber("123")
	var _ types.LabResult = lab.MakeResult("", "", phone, false)
}

func TestResultGetter(t *testing.T) {
	phone := types.MakePhoneNumber("123")
	result := lab.MakeResult("abc", "bob c lamptest", phone, false)

	assert.Equal(t, "abc", result.Id())
	assert.Equal(t, "bob c lamptest", result.FullName())
	assert.Equal(t, phone, result.CellPhoneNumber())
	assert.False(t, result.IsPositive())
}
