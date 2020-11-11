package pipe_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/comptag/bobcat-lamp/internal/pipe"
	"github.com/comptag/bobcat-lamp/internal/types"
)

func TestRawPatientIsValid(t *testing.T) {
	r := pipe.RawPatient{}
	assert.False(t, r.IsValid())

	r = pipe.RawPatient{"a", "b", ""}
	assert.False(t, r.IsValid())

	r = pipe.RawPatient{"a", "", "c"}
	assert.False(t, r.IsValid())

	r = pipe.RawPatient{"a", "b", ""}
	assert.False(t, r.IsValid())

	r = pipe.RawPatient{"a", "b", "c"}
	assert.True(t, r.IsValid())
}

func TestRawLabResultIsValid(t *testing.T) {
	r := pipe.RawLabResult{}
	assert.False(t, r.IsValid())

	r = pipe.RawLabResult{"a", ""}
	assert.False(t, r.IsValid())

	r = pipe.RawLabResult{"", "b"}
	assert.False(t, r.IsValid())

	r = pipe.RawLabResult{"x", "b"}
	assert.False(t, r.IsValid())

	r = pipe.RawLabResult{"x", "positive"}
	assert.True(t, r.IsValid())

	r = pipe.RawLabResult{"x", "negative"}
	assert.True(t, r.IsValid())
}

func TestRawLabResultToLabResult(t *testing.T) {
	alice := types.MakePatient("123", "alice a", types.MakePhoneNumber("4065551234"))

	r, err := pipe.RawLabResult{"123", "positive"}.ToLabResult(alice)
	assert.NoError(t, err)
	assert.Equal(t, alice, r.Patient())
	assert.True(t, true, r.IsPositive())

	r, err = pipe.RawLabResult{"123", "negative"}.ToLabResult(alice)
	assert.NoError(t, err)
	assert.Equal(t, alice, r.Patient())
	assert.True(t, true, r.IsPositive())

	r, err = pipe.RawLabResult{"123", "native"}.ToLabResult(alice)
	assert.Error(t, err)

	r, err = pipe.RawLabResult{"12", "negative"}.ToLabResult(alice)
	assert.Error(t, err)
}

func TestLoadPatients(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		patientsReader := strings.NewReader(`student_id,full_name,cell_phone_number
123,alice a,4065551234
124,bob b,4065551235`)

		patients, err := pipe.LoadPatients(patientsReader)
		assert.NotNil(t, patients)
		assert.NoError(t, err)

		alice := types.MakePatient("123", "alice a", types.MakePhoneNumber("4065551234"))
		assert.Equal(t, alice, patients[0])

		bob := types.MakePatient("124", "bob b", types.MakePhoneNumber("4065551235"))
		assert.Equal(t, bob, patients[1])
	})

	t.Run("missing field", func(t *testing.T) {
		patientsReader := strings.NewReader(`student_id,cell_phone_number
123,4065551234
124,4065551235`)

		patients, err := pipe.LoadPatients(patientsReader)
		assert.Nil(t, patients)
		assert.Error(t, err)
	})
}

func TestLoadResults(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		resultsReader := strings.NewReader(`student_id,result
123,positive
124,negative`)

		results, err := pipe.LoadLabResults(resultsReader)
		assert.NotNil(t, results)
		assert.NoError(t, err)

		assert.Equal(t, pipe.RawLabResult{"123", "positive"}, *results[0])
		assert.Equal(t, pipe.RawLabResult{"124", "negative"}, *results[1])
	})

	t.Run("missing field", func(t *testing.T) {
		resultsReader := strings.NewReader(`student_id
123
124`)

		results, err := pipe.LoadLabResults(resultsReader)
		assert.Nil(t, results)
		assert.Error(t, err)
	})

	t.Run("typo", func(t *testing.T) {
		resultsReader := strings.NewReader(`student_id,result
123,pasitive
124,negative`)

		results, err := pipe.LoadLabResults(resultsReader)
		assert.Nil(t, results)
		assert.Error(t, err)
	})
}

func TestHashJoin(t *testing.T) {
	aliceResult := pipe.RawLabResult{"123", "positive"}
	bobResult := pipe.RawLabResult{"124", "negative"}

	alice := types.MakePatient("123", "alice a", types.MakePhoneNumber("4065551234"))
	bob := types.MakePatient("124", "bob b", types.MakePhoneNumber("4065551235"))

	t.Run("happy path", func(t *testing.T) {
		rawResults := []*pipe.RawLabResult{&aliceResult, &bobResult}
		patients := []types.Patient{bob, alice}

		results, err := pipe.HashJoin(rawResults, patients)
		assert.NoError(t, err)

		assert.Equal(t, bob, results[0].Patient())
		assert.Equal(t, false, results[0].IsPositive())

		assert.Equal(t, alice, results[1].Patient())
		assert.Equal(t, true, results[1].IsPositive())
	})

	t.Run("patient but no result", func(t *testing.T) {
		rawResults := []*pipe.RawLabResult{&aliceResult}
		patients := []types.Patient{bob, alice}

		_, err := pipe.HashJoin(rawResults, patients)
		assert.Error(t, err)
	})

	t.Run("result but no patient", func(t *testing.T) {
		rawResults := []*pipe.RawLabResult{&aliceResult, &bobResult}
		patients := []types.Patient{alice}

		_, err := pipe.HashJoin(rawResults, patients)
		assert.Error(t, err)

	})

	t.Run("result key is not unique", func(t *testing.T) {
		rawResults := []*pipe.RawLabResult{&aliceResult, &bobResult, &aliceResult}
		patients := []types.Patient{bob, alice}

		_, err := pipe.HashJoin(rawResults, patients)
		assert.Error(t, err)
	})

	t.Run("patient key is not unique", func(t *testing.T) {
		rawResults := []*pipe.RawLabResult{&aliceResult, &bobResult}
		patients := []types.Patient{bob, alice, bob}

		_, err := pipe.HashJoin(rawResults, patients)
		assert.Error(t, err)
	})
}

func TestLoad(t *testing.T) {

	alice := types.MakePatient("123", "alice a", types.MakePhoneNumber("4065551234"))
	bob := types.MakePatient("124", "bob b", types.MakePhoneNumber("4065551235"))

	patientsReader := strings.NewReader(`student_id,full_name,cell_phone_number
123,alice a,4065551234
124,bob b,4065551235`)

	resultsReader := strings.NewReader(`student_id,result
123,positive
124,negative`)

	results, err := pipe.Load(patientsReader, resultsReader)

	assert.NotNil(t, results)
	assert.NoError(t, err)

	assert.Equal(t, alice, results[0].Patient())
	assert.Equal(t, true, results[0].IsPositive())
	assert.Equal(t, bob, results[1].Patient())
	assert.Equal(t, false, results[1].IsPositive())
}

func TestLoad(t *testing.T) {

	alice := types.MakePatient("123", "alice a", types.MakePhoneNumber("4065551234"))
	bob := types.MakePatient("124", "bob b", types.MakePhoneNumber("4065551235"))

	patientsReader := strings.NewReader(`student_id,full_name,cell_phone_number
123,alice a,4065551234
124,bob b,4065551235`)

	resultsReader := strings.NewReader(`student_id,result
123,positive
124,negative`)

	results, err := pipe.Load(patientsReader, resultsReader)

	assert.NotNil(t, results)
	assert.NoError(t, err)

	assert.Equal(t, alice, results[0].Patient())
	assert.Equal(t, true, results[0].IsPositive())
	assert.Equal(t, bob, results[1].Patient())
	assert.Equal(t, false, results[1].IsPositive())
}
