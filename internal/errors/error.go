package errors

import "errors"

var (
	ErrCredentialError          = errors.New("credential error")
	ErrDuplicateEmailError      = errors.New("this email address is already registered")
	ErrOldPasswordNotMatchError = errors.New("old password does not match")
	ErrTopUpTooSmallError       = errors.New("top up amount is too small (minimum Rp. 10,000)")
	ErrInvalidRoleAddHouseError = errors.New("only host can add a new house")
)
