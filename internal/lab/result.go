package lab

type Result struct {
	id         string
	name       string
	number     string
	testResult bool
}

func MakeResult(id, name, number string, result bool) Result {
	return Result{id, "", "", true}

}

func (r Result) Id() string {
	return r.id
}

//
// func (r Result) FullName() string {
// }
//
// func (r Result) CellPhoneNumber() string {
// }
//
// func (r Result) IsPositive() bool {
// }
