package response

import (
	"errors"
	"net/http"
)

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
	ErrPasswordNotMatch      = errors.New("password not match")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral    = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest = NewError("bad request", "40000", http.StatusBadRequest)
)

var (
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorAuthIsNotExist        = NewError(ErrAuthIsNotExist.Error(), "40401", http.StatusNotFound)
	ErrorEmailAlreadyUsed      = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch      = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
	ErrorNotFound              = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
)

var (
	ErrorMapping = map[string]Error{
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailRequired,
		ErrPasswordRequired.Error():      ErrorEmailRequired,
		ErrPasswordInvalidLength.Error(): ErrorEmailRequired,
		ErrAuthIsNotExist.Error():        ErrorEmailRequired,
		ErrEmailAlreadyUsed.Error():      ErrorEmailRequired,
		ErrPasswordNotMatch.Error():      ErrorEmailRequired,
		ErrNotFound.Error():              ErrorNotFound,
	}
)
