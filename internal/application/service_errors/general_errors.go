package service_errors

import "net/http"

var ErrIdNotValid = ServiceError{
	HttpStatus: http.StatusBadRequest,
	Message:    "Invalid id",
	Code:       "INVALID_ID",
}
