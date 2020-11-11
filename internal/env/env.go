package env

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/comptag/bobcat-lamp/internal/types"
	"github.com/joho/godotenv"
)

type TwilioConfig struct {
	AccountSid          string
	AuthToken           string
	FromNumber          types.PhoneNumber
	PollIntervalSeconds time.Duration
	MaxRetries          int
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Error reading config file", err)
	}
	return i

}

func LoadEnv() TwilioConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return TwilioConfig{
		os.Getenv("TWILIO_ACCOUNT_SID"),
		os.Getenv("TWILIO_AUTH_TOKEN"),
		types.MakePhoneNumber(os.Getenv("TWILIO_FROM_NUMBER")),
		time.Duration(toInt(os.Getenv("TWILIO_POLL_INTERVAL_SECONDS"))),
		toInt(os.Getenv("TWILIO_MAT_RETRIES")),
	}
}
