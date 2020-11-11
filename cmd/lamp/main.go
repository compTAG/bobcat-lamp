package main

import (
	"fmt"
	"time"

	"github.com/comptag/bobcat-lamp/internal/env"
	"github.com/comptag/bobcat-lamp/internal/lab"
	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/internal/types"
)

const LiveSms = true

func main() {
	twilioCfg := env.LoadEnv()

	cell := types.MakePhoneNumber("9196271828")
	me := types.MakePatient("daveID", "David Millman", cell)
	result := lab.MakeResult(me, false)

	reporter := msg.MakeDummyReporter()
	if LiveSms {
		reporter = msg.MakeSmsReporter(
			twilioCfg.AccountSid,
			twilioCfg.AuthToken,
			twilioCfg.FromNumber,
			twilioCfg.PollIntervalSeconds*time.Second,
			twilioCfg.MaxRetries,
		)
	}

	r, err := reporter.Report(result)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", r)
	}
}
