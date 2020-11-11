package main

import (
	"fmt"
	"time"

	"github.com/comptag/bobcat-lamp/internal/env"
	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/internal/pipe"
)

const LiveSms = false

func main() {
	twilioCfg := env.LoadEnv()

	// setup the resporer
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

	// load results
	results, err := pipe.LoadFile(
		"./testdata/patients.csv",
		"./testdata/results.csv",
	)
	if err != nil {
		fmt.Println("Error", err)
	}

	// send a report
	for _, result := range results {
		r, err := reporter.Report(result)

		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("Success", r)
		}
	}
}
