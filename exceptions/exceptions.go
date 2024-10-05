package exceptions

import "errors"

var (
	ErrInvalidToken          = errors.New("provided token is invalid")
	ErrInvalidSigningMethod  = errors.New("provided token signed incorrectly")
	ErrBadRequest            = errors.New("provided request body or parametes is invalid")
	ErrNotFound              = errors.New("such item was not found")
	ErrInternalServerError   = errors.New("internal server error")
	ErrPasswordIncorrect     = errors.New("provided password is incorrect")
	ErrUserDoesNotExists     = errors.New("such user does not exists")
	ErrSuchUserAlreadyExists = errors.New("such user already exists")
)
