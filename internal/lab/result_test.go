package lab_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/comptag/bobcat-lamp/internal/lab"
	"github.com/comptag/bobcat-lamp/internal/types"
)

func TestResultInterface(t *testing.T) {
	// if result is not a lab result this code will cause compilation to fail
	patient := types.Patient{}
	var _ types.LabResult = lab.MakeResult(patient, false)
}

func TestResultGetter(t *testing.T) {
	phone := types.MakePhoneNumber("123")
	patient := types.MakePatient("id", "bob c lamptest", phone)
	result := lab.MakeResult(patient, false)

	assert.Equal(t, patient, result.Patient())
	assert.False(t, result.IsPositive())
}
