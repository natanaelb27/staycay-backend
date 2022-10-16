package errors

import "errors"

var (
	ErrCredentialError                  = errors.New("credential error")
	ErrDuplicateEmailError              = errors.New("this email address is already registered")
	ErrOldPasswordNotMatchError         = errors.New("old password does not match")
	ErrTopUpTooSmallError               = errors.New("top up amount is too small (minimum Rp. 10,000)")
	ErrInvalidRoleNotHostError          = errors.New("user is not a host")
	ErrInvalidHostError                 = errors.New("user is not the host")
	ErrInvalidUserReservationError      = errors.New("user is not the customer")
	ErrDeleteUpcomingReservationError   = errors.New("house cannot be deleted because there is an upcoming reservation, house status will be changed to inactive instead")
	ErrIdNotFoundError                  = errors.New("this id is not found")
	ErrNotValidReservationDateConflict  = errors.New("reservation is not valid due to conflicting date")
	ErrInsufficientBalanceError         = errors.New("insufficient balance")
	ErrNotTodaysReservationError        = errors.New("this is not today's reservation")
	ErrInvalidReservationDateLaterError = errors.New("invalid reservation date, the check in date is later than the check out date")
	ErrInvalidReservationDateSameError  = errors.New("invalid reservation date, the check in date is the same as the check out date")
	ErrInvalidReservationBeforeNowError = errors.New("invalid reservation date, the reservation date is earlier than today")
	ErrDuplicatePickupReservationError  = errors.New("this reservation is already registered for pickup")
	ErrPickupCompletedReservation       = errors.New("pickup is not available for completed reservation")
	ErrPickupOngoingReservation         = errors.New("pickup is not available for ongoing reservation")
	ErrInsufficientGameChanceError      = errors.New("insufficient game chance")
	// ErrInvalidRoleAddHouseError = errors.New("only host can add a new house")
)
