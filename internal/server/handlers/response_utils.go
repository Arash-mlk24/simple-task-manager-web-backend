package handlers

import "github.com/Arash-mlk24/simple-task-manager-web-backend/internal/application/dto"

func ApiSuccess[T any](result *T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		HasError: false,
		Result:   result,
	}
}

func ApiFailure(code, message string) dto.ApiResponse[any] {
	return dto.ApiResponse[any]{
		HasError: true,
		Error: &dto.ApiError{
			Code:    code,
			Message: message,
		},
	}
}
