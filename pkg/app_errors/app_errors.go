package app_errors

import (
	"net/http"

	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/pkg/errors"
)

var (
	// Http Error Handler
	ErrRouteNotFound         = errors.New("route not found")
	ErrMalformatBody         = errors.New("malformed_body")
	ErrValidation            = errors.New("validation_error")
	ErrInvalidMultiPart      = errors.New("invalid multipart content-type")
	ErrMissingOrMalformatJWT = errors.New("Missing or malformed JWT")

	ErrNotFound            = errors.New("not found")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrForbiddenAccess     = errors.New("forbidden access")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrConflict            = errors.New("conflict")
	ErrTimeout             = errors.New("timeout")
	ErrBadRequest          = errors.New("bad request")
	ErrGateway             = errors.New("gateway")

	//Guards
	ErrMissingJWTToken  = errors.New("Missing JWT token")
	ErrInvalidJWTToken  = errors.New("Invalid JWT token")
	ErrJWTTokenNotValid = errors.New("JWT token is not valid")
	ErrAccessDenied     = errors.New("Access Denied")

	ErrInvalidRole = errors.New("invalid role")
)

type AppError struct {
	Code    int
	Err     error
	Message constants.ResponseMessage
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (h AppError) Error() string {
	return h.Err.Error()
}

func BadRequest(err error) error {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: constants.BadRequest,
		Err:     err,
	}
}

func InternalServerError(err error) error {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: constants.InternalServerError,
		Err:     err,
	}
}

func Unauthorized(err error) error {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: constants.Unauthorized,
		Err:     err,
	}
}

func Forbidden(err error) error {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: constants.Forbidden,
		Err:     err,
	}
}

func NotFound(err error) error {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: constants.NotFound,
		Err:     err,
	}
}

func Conflict(err error) error {
	return &AppError{
		Code:    http.StatusConflict,
		Message: constants.Conflict,
		Err:     err,
	}
}

func GatewayTimeout(err error) error {
	return &AppError{
		Code:    http.StatusGatewayTimeout,
		Message: constants.GatewayTimeOut,
		Err:     err,
	}
}
