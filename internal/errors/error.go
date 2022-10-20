package errors

import "errors"

var (
	ErrCredentialError                           = errors.New("credential error")
	ErrMissingFieldError                         = errors.New("field is missing")
	ErrInvalidEmailInputError                    = errors.New("email is not valid")
	ErrInvalidPasswordInputError                 = errors.New("password does not meet the requirements")
	ErrDuplicateEmailError                       = errors.New("this email address is already registered")
	ErrOldPasswordNotMatchError                  = errors.New("old password does not match")
	ErrTopUpTooSmallError                        = errors.New("top up amount is too small (minimum Rp. 10,000)")
	ErrInvalidRoleNotHostError                   = errors.New("user is not a host")
	ErrInvalidRoleNotAdminError                  = errors.New("user is not an admin")
	ErrInvalidHostError                          = errors.New("user is not the host")
	ErrInvalidUserReservationError               = errors.New("user is not the customer")
	ErrInvalidHostAndUserReservationError        = errors.New("user is neither the host or the customer")
	ErrDeleteUpcomingReservationError            = errors.New("house cannot be deleted because there is an upcoming reservation, house status will be changed to inactive instead")
	ErrIdNotFoundError                           = errors.New("this id is not found")
	ErrNotValidReservationDateConflict           = errors.New("reservation is not valid due to conflicting date")
	ErrInsufficientBalanceError                  = errors.New("insufficient balance")
	ErrNotTodaysReservationError                 = errors.New("this is not today's reservation")
	ErrInvalidReservationDateLaterError          = errors.New("invalid reservation date, the check in date is later than the check out date")
	ErrInvalidReservationDateSameError           = errors.New("invalid reservation date, the check in date is the same as the check out date")
	ErrInvalidReservationBeforeNowError          = errors.New("invalid reservation date, the reservation date is earlier than today")
	ErrDuplicatePickupReservationError           = errors.New("this reservation is already registered for pickup")
	ErrPickupCompletedReservation                = errors.New("pickup is not available for completed reservation")
	ErrPickupOngoingReservation                  = errors.New("pickup is not available for ongoing reservation")
	ErrInsufficientGameChanceError               = errors.New("insufficient game chance")
	ErrInvalidCheckInDateUpdatePickupStatusError = errors.New("cannot update pickup status, check in date is not today")
	// ErrInvalidRoleAddHouseError = errors.New("only host can add a new house")
)

// type MissingField struct {
// 	field string
// }

// func (err MissingField) MissingFieldError() string {
// 	return err.field + " field is mising"
// }
