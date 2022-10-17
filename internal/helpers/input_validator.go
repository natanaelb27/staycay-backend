package helpers

import (
	"net/mail"
	"unicode"
)

func VerifyEMailAddress(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func VerifyPassword(password string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			return false
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber
}
