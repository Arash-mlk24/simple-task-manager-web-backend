package dto

type ApiResponse[T any] struct {
	HasError bool      `json:"hasError"`
	Result   *T        `json:"result,omitempty"`
	Error    *ApiError `json:"error,omitempty"`
}

type ApiError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type DocsApiResponse struct {
	HasError bool     `json:"hasError"`
	Result   any      `json:"result,omitempty"`
	Error    ApiError `json:"error,omitempty"`
}
