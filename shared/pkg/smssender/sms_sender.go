package sendersms

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SenderSMS struct {
	number                  string
	autorizationHeaderValue string
	client                  *resty.Client
}

func NewSenderSMS(number string, apiKey string, client *resty.Client) *SenderSMS {
	if client == nil {
		client = resty.New()
	}
	autorizationHeaderValue := fmt.Sprintf("Bearer %s", apiKey)
	return &SenderSMS{
		number:                  number,
		autorizationHeaderValue: autorizationHeaderValue,
		client:                  client,
	}
}

func (sender *SenderSMS) Send(destination string, text string) error {
	json := struct {
		Number      string `json:"number"`
		Destination string `json:"destination"`
		Text        string `json:"text"`
	}{
		Number:      sender.number,
		Destination: destination,
		Text:        text,
	}
	req := sender.client.
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(json)
	res, err := req.Post("https://api.exolve.ru/messaging/v1/SendSMS")
	if err != nil {
		return err
	}
	if res.StatusCode() > 399 {
		return fmt.Errorf("SenderSMS error: %s", res.String())
	}
	return nil
}

type SenderSMSMock struct {
}

func NewSenderSMSMock() *SenderSMSMock {
	return &SenderSMSMock{}
}

func (sender *SenderSMSMock) Send(destination string, text string) error {
	fmt.Println(text)
	return nil
}
