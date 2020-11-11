package main

import (
	"time"

	"github.com/labstack/gommon/log"

	"github.com/comptag/bobcat-lamp/internal/env"
	"github.com/comptag/bobcat-lamp/internal/lab"
	"github.com/comptag/bobcat-lamp/internal/msg"
	"github.com/comptag/bobcat-lamp/internal/pipe"
)

const LiveSms = true

func getFilenames() (string, string, error) {
	return "./testdata/patients.csv", "./testdata/results.csv", nil

}

func initReporter(useLiveSMS bool) msg.Reporter {
	twilioCfg := env.LoadEnv()

	reporter := msg.MakeDummyReporter()
	if useLiveSMS {
		reporter = msg.MakeSmsReporter(
			twilioCfg.AccountSid,
			twilioCfg.AuthToken,
			twilioCfg.FromNumber,
			twilioCfg.PollIntervalSeconds*time.Second,
			twilioCfg.MaxRetries,
		)
	}

	return reporter
}

func sendReport(reporter msg.Reporter, results []*lab.Result) ([]string, error) {
	sids := make([]string, len(results))
	for i, result := range results {
		r, err := reporter.Report(result)

		if err != nil {
			return sids, err
		}

		sids[i] = r
	}
	return sids, nil
}

func main() {
	// init variables
	reporter := initReporter(LiveSms)
	patientsFileName, resultsFileName, err := getFilenames()
	if err != nil {
		log.Error(err)
		return
	}

	// load results
	results, err := pipe.LoadFile(patientsFileName, resultsFileName)
	if err != nil {
		log.Error(err)
		return
	}

	// send a report
	sids, err := sendReport(reporter, results)
	if err != nil {
		log.Error(err)
		return
	}

	log.Printf("Success %v", sids)
}
