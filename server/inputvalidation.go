package main

import (
	"regexp"

	gocheckpassword "github.com/ninja-software/go-check-passwd"
)

func validateEmail(email string) error {
	result := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !result.MatchString(email) {
		return InvalidEmail
	}
	return nil
}

func validatePassword(password string) error {

	if len(password) >= 8 && !gocheckpassword.IsCommon(password) {
		return nil
	}

	return IncorrectPasswordFormat
}
