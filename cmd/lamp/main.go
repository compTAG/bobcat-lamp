package main

import (
	"fmt"
	"time"

	"github.com/comptag/bobcat-lamp/internal/lab"
	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/internal/types"
)

const LiveSms = true

// move to config
const TwilioAccountSid = "***"
const TwilioAuthToken = "***"
const TwilioFromNumber = types.MakePhoneNumber("4065511606")
const PollInterval = time.Second
const MaxTries = 5

func main() {

	cell := types.MakePhoneNumber("9196271828")
	me := types.MakePatient("daveID", "David Millman", cell)
	result := lab.MakeResult(me, false)

	reporter := msg.MakeDummyReporter()
	if LiveSms {
		reporter = msg.MakeSmsReporter(
			TwilioAccountSid,
			TwilioAuthToken,
			TwilioFromNumber,
			PollInterval,
			MaxTries,
		)
	}

	r, err := reporter.Report(result)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", r)
	}
}
