package types

type PhoneNumber string

func MakePhoneNumber(p string) PhoneNumber {
	return PhoneNumber(p)
}

type LabResult interface {
	Patient() Patient
	IsPositive() bool
}
