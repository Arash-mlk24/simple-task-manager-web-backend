package service_errors

import "net/http"

var ErrUnauthorized = ServiceError{
	HttpStatus: http.StatusUnauthorized,
	Message:    "unauthorized",
	Code:       "UNAUTHORIZED",
}

var ErrForbidden = ServiceError{
	HttpStatus: http.StatusForbidden,
	Message:    "forbidden",
	Code:       "FORBIDDEN",
}

var ErrEmailOrPasswordMismatch = ServiceError{
	HttpStatus: http.StatusUnauthorized,
	Message:    "email or password mismatch",
	Code:       "EMAIL_OR_PASSWORD_MISMATCH",
}
