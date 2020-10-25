package main

import (
	"fmt"

	"github.com/comptag/bobcat-lamp/internal/lab"
	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/internal/types"
)

func main() {

	cell := types.MakePhoneNumber("5555555555")
	jane := types.MakePatient("a-student-id", "Jane Doe", cell)
	result := lab.MakeResult(jane, true)

	reporter := msg.MakeDummyReporter()
	r, err := reporter.Report(result)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", r)
	}
}
