package smssender

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ISMSSender interface {
	Send(destination string, text string) error
}

type SMSSender struct {
	number                  string
	autorizationHeaderValue string
	client                  *resty.Client
}

func New(number string, apiKey string, client *resty.Client) *SMSSender {
	if client == nil {
		client = resty.New()
	}
	autorizationHeaderValue := fmt.Sprintf("Bearer %s", apiKey)
	return &SMSSender{
		number:                  number,
		autorizationHeaderValue: autorizationHeaderValue,
		client:                  client,
	}
}

func (sender *SMSSender) Send(destination string, text string) error {
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
		SetHeader("Authorization", sender.autorizationHeaderValue).
		SetBody(json)
	resp, err := req.Post("https://api.exolve.ru/messaging/v1/SendSMS")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("SMSSender error: %s", resp.String())
	}
	return nil
}

type SMSSenderMock struct {
}

func NewSenderSMSMock() *SMSSenderMock {
	return &SMSSenderMock{}
}

func (sender *SMSSenderMock) Send(destination string, text string) error {
	fmt.Println(text)
	return nil
}
