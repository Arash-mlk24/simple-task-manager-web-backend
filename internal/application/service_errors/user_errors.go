package service_errors

import "net/http"

var ErrUserNotFound = ServiceError{
	HttpStatus: http.StatusNotFound,
	Message:    "User not found",
	Code:       "USER_NOT_FOUND",
}
