package main

import (
	"context"
	"fmt"

	"github.com/mailgun/mailgun-go/v3"
)

//verify email after signup
func (e *SendEmailInfo) VerifyEmail(email string, veriToken string) error {
	url := fmt.Sprintf("%s://%s/api/confirm-email/%s", e.Scheme, e.ServerDomain, veriToken)
	mg := mailgun.NewMailgun(e.EmailDomain, e.EmailAPIKey)
	m := mg.NewMessage(
		"Crypto Tracker<mailgun@"+e.EmailDomain+">",
		"Confirm Your Registration",
		"You're one click away from getting latest information on cryptocurrencies! \n\nPlease click on the link below to verify your account: \n"+url,
		email,
	)
	_, _, err := mg.Send(context.Background(), m)
	if err != nil {
		return VerifyEmailError
	}
	return nil
}

//reset password for Forgot Password
func (e *SendEmailInfo) ResetPassword(email string, resetPassToken string) error {
	url := fmt.Sprintf("%s://%s/api/reset-password/%s", e.Scheme, e.ServerDomain, resetPassToken)
	mg := mailgun.NewMailgun(e.EmailDomain, e.EmailAPIKey)
	m := mg.NewMessage(
		"Crypto Tracker<mailgun@"+e.EmailDomain+">",
		"Reset Your Password",
		"You've requested a password reset.\n To change your password, click on the link below: \n"+url,
		email,
	)
	_, _, err := mg.Send(context.Background(), m)
	if err != nil {
		return ResetPasswordError
	}
	return nil
}
