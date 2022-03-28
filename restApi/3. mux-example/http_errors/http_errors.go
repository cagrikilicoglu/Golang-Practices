package http_errors

import (
	"errors"
	"net/http"
)

var (
	ContentTypeError    = errors.New("Content-Type must be `application/json`")
	InternalServerError = errors.New("Internal Server Error")
	CannotDecodeError   = errors.New("Payload cannot be decoded")
)

type RestError struct {
	ErrStatus int    `json:"code",omitempty`
	ErrError  string `json:"message",omitempty`
}

func NewRestError(status int, err string) RestError {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
	}
}

func NewInternalServerError() RestError {
	return RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  InternalServerError.Error(),
	}
}

// func ParseError(err error) RestError {

// }
