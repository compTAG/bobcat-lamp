package types

type PhoneNumber struct {
	string
}

func MakePhoneNumber(p string) PhoneNumber {
	return PhoneNumber{p}
}

func (p PhoneNumber) InternationalNoDash() string {
	return p.string
}

type LabResult interface {
	Patient() Patient
	IsPositive() bool
}
