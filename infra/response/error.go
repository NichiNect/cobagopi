package response

import "errors"

// ? General error
var (
	ErrNotFound = errors.New("not found")
)

// ? Auth error
var (
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimum 6 characters")
	ErrAuthIsNotExist        = errors.New("auth is not exists")
	ErrEmailAlreadyUsed      = errors.New("email already used")
)
