package service_errors

type ServiceError struct {
	HttpStatus int
	Message    string
	Code       string
}
