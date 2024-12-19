package response

import (
	"errors"
	"net/http"
)

// ? General error
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)

// ? Auth error
var (
	// auth
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimum 6 characters")
	ErrAuthIsNotExist        = errors.New("auth is not exists")
	ErrEmailAlreadyUsed      = errors.New("email already used")
	ErrPasswordNotMatch      = errors.New("password not match")

	// product
	ErrProductRequired = errors.New("product is required")
	ErrProductInvalid  = errors.New("product must have minimum 4 characters")
	ErrStockRequired   = errors.New("product stock is required")
	ErrStockInvalid    = errors.New("product stock must greater than 0")
	ErrPriceRequired   = errors.New("product price is required")
	ErrPriceInvalid    = errors.New("product price must greater than 0")

	// transaction
	ErrAmountInvalid          = errors.New("invalid amount")
	ErrAmountGreaterThanStock = errors.New("amount greater than stock")
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
	ErrorGeneral      = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest   = NewError("bad request", "40000", http.StatusBadRequest)
	ErrorNotFound     = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
	ErrorUnauthorized = NewError(ErrUnauthorized.Error(), "40100", http.StatusUnauthorized)
)

var (
	// error bad request
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)

	ErrorProductRequired = NewError(ErrProductRequired.Error(), "40005", http.StatusBadRequest)
	ErrorProductInvalid  = NewError(ErrProductInvalid.Error(), "40006", http.StatusBadRequest)
	ErrorStockRequired   = NewError(ErrStockRequired.Error(), "40007", http.StatusBadRequest)
	ErrorStockInvalid    = NewError(ErrStockInvalid.Error(), "40008", http.StatusBadRequest)
	ErrorPriceRequired   = NewError(ErrPriceRequired.Error(), "40009", http.StatusBadRequest)
	ErrorPriceInvalid    = NewError(ErrPriceInvalid.Error(), "40010", http.StatusBadRequest)
	ErrorInvalidAmount   = NewError(ErrAmountInvalid.Error(), "40011", http.StatusBadRequest)

	ErrorAuthIsNotExist   = NewError(ErrAuthIsNotExist.Error(), "40401", http.StatusNotFound)
	ErrorEmailAlreadyUsed = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
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
		ErrUnauthorized.Error():          ErrorUnauthorized,
	}
)
