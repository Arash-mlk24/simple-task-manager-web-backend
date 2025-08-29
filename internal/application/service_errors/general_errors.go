package service_errors

import "net/http"

var ErrIdNotValid = ServiceError{
	HttpStatus: http.StatusBadRequest,
	Message:    "Invalid id",
	Code:       "INVALID_ID",
}

var ErrInvalidData = ServiceError{
	HttpStatus: http.StatusBadRequest,
	Message:    "Invalid data",
	Code:       "INVALID_DATA",
}

var ErrInternal = ServiceError{
	HttpStatus: http.StatusInternalServerError,
	Message:    "Internal server error",
	Code:       "INTERNAL_SERVER_ERROR",
}
