package types

type Patient struct {
	id     string
	name   string
	number PhoneNumber
}

func MakePatient(id, name string, number PhoneNumber) Patient {
	return Patient{id, name, number}
}

func (r Patient) Id() string {
	return r.id
}

func (r Patient) FullName() string {
	return r.name
}

func (r Patient) CellPhoneNumber() PhoneNumber {
	return r.number
}
