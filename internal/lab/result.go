package lab

import "github.com/comptag/bobcat-lamp/internal/types"

type Result struct {
	patient    types.Patient
	testResult bool
}

func MakeResult(patient types.Patient, result bool) Result {
	return Result{patient, result}

}

func (r Result) Patient() types.Patient {
	return r.patient
}

func (r Result) IsPositive() bool {
	return r.testResult
}
