package errors

import "errors"

var (
	ErrCredentialError                  = errors.New("credential error")
	ErrDuplicateEmailError              = errors.New("this email address is already registered")
	ErrOldPasswordNotMatchError         = errors.New("old password does not match")
	ErrTopUpTooSmallError               = errors.New("top up amount is too small (minimum Rp. 10,000)")
	ErrInvalidRoleNotHostError          = errors.New("user is not a host")
	ErrInvalidHostError                 = errors.New("user is not the host")
	ErrDeleteUpcomingReservationError   = errors.New("house cannot be deleted because there is an upcoming reservation, house status will be changed to inactive instead")
	ErrIdNotFoundError                  = errors.New("this id is not found")
	ErrNotValidReservationDateConflict  = errors.New("reservation is not valid due to conflicting date")
	ErrBalanceNotEnoughError            = errors.New("user balance is not enough")
	ErrNotTodaysReservationError        = errors.New("this is not today's reservation")
	ErrInvalidReservationDateLaterError = errors.New("invalid reservation date, the check in date is later than the check out date")
	ErrInvalidReservationDateSameError  = errors.New("invalid reservation date, the check in date is the same as the check out date")
	ErrInvalidReservationBeforeNowError = errors.New("invalid reservation date, the reservation date is earlier than today")
	// ErrInvalidRoleAddHouseError = errors.New("only host can add a new house")
)
