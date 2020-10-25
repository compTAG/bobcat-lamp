package types_test

import (
	"testing"

	"github.com/comptag/bobcat-lamp/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestPatientGetter(t *testing.T) {
	phone := types.MakePhoneNumber("123")

	patient := types.MakePatient("id", "bob z", phone)
	assert.Equal(t, "id", patient.Id())
	assert.Equal(t, "bob z", patient.FullName())
	assert.Equal(t, phone, patient.CellPhoneNumber())
}
