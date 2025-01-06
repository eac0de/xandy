package smssender

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ISMSSender interface {
	Send(destination string, text string) error
}

type SMSSender struct {
	number string
	token  string
	client *resty.Client
}

func New(number string, token string) *SMSSender {
	return &SMSSender{
		number: number,
		token:  token,
		client: resty.New(),
	}
}

func (s *SMSSender) Send(destination string, text string) error {
	json := struct {
		Number      string `json:"number"`
		Destination string `json:"destination"`
		Text        string `json:"text"`
	}{
		Number:      s.number,
		Destination: destination,
		Text:        text,
	}
	resp, err := s.client.
		R().
		SetAuthToken(s.token).
		SetBody(json).
		Post("https://api.exolve.ru/messaging/v1/SendSMS")
	if err != nil {
		return err
	}
	if resp.StatusCode() > 299 {
		return fmt.Errorf("SMSSender error: %s", resp.String())
	}
	return nil
}

type SMSSenderMock struct {
}

func NewMock() *SMSSenderMock {
	return &SMSSenderMock{}
}

func (s *SMSSenderMock) Send(destination string, text string) error {
	fmt.Println(text)
	return nil
}
