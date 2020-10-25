package lab

import "github.com/comptag/bobcat-lamp/internal/types"

type Result struct {
	id         string
	name       string
	number     types.PhoneNumber
	testResult bool
}

func MakeResult(id, name string, number types.PhoneNumber, result bool) Result {
	return Result{id, name, number, result}

}

func (r Result) Id() string {
	return r.id
}

func (r Result) FullName() string {
	return r.name
}

func (r Result) CellPhoneNumber() types.PhoneNumber {
	return r.number
}

func (r Result) IsPositive() bool {
	return r.testResult
}
