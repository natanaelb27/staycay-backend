package errors

import "errors"

var (
	ErrCredentialError     = errors.New("credential error")
	ErrDuplicateEmailError = errors.New("this email address is already registered")
)
