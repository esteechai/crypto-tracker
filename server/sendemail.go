package main

import (
	"context"
	"fmt"

	"github.com/mailgun/mailgun-go/v3"
)

func (e *SendEmailInfo) VerifyEmail(email string, veriToken string) error {
	url := fmt.Sprintf("%s://%s/api/confirm-email/%s", e.Scheme, e.ServerDomain, veriToken)
	mg := mailgun.NewMailgun(e.EmailDomain, e.EmailAPIKey)
	m := mg.NewMessage(
		"Crypto Tracker<mailgun@"+e.EmailDomain+">",
		"Confirm Your Registration",
		"You're one click away from getting latest information on cryptocurrencies! \n\nPlease click on the link below to verify your account: \n"+url,
		"esteechai@gmail.com",
	)
	_, _, err := mg.Send(context.Background(), m)
	if err != nil {
		return VerifyEmailError
	}
	return nil
}
