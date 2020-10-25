package types

type PhoneNumber string

func MakePhoneNumber(p string) PhoneNumber {
	return PhoneNumber(p)
}

type LabResult interface {
	Id() string
	FullName() string
	CellPhoneNumber() PhoneNumber
	IsPositive() bool
}
