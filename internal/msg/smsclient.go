package msg

import (
	"errors"

	"github.com/comptag/bobcat-lamp/internal/types"
	"github.com/sfreiberg/gotwilio"
)

type SmsResponse = gotwilio.SmsResponse

type SmsClient interface {
	Send(from, to types.PhoneNumber, body string) (*SmsResponse, error)
	Get(sid string) (*SmsResponse, error)
}

type TwilioClient struct {
	twilio *gotwilio.Twilio
}

func MakeTwilioClient(accountSid, authToken string) TwilioClient {
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)
	return TwilioClient{twilio}
}

func mergeError(excep *gotwilio.Exception, err error) error {
	var result error = nil
	if err != nil {
		result = err
	} else if excep != nil {
		result = errors.New(excep.Error())
	}
	return result
}

func (c TwilioClient) Send(from, to types.PhoneNumber, body string) (*SmsResponse, error) {
	response, excep, err := c.twilio.SendSMS(
		from.InternationalNoDash(),
		to.InternationalNoDash(),
		body,
		"",
		"",
	)
	return response, mergeError(excep, err)
}

func (c TwilioClient) Get(sid string) (*SmsResponse, error) {
	response, excep, err := c.twilio.GetSMS(sid)
	return response, mergeError(excep, err)
}
