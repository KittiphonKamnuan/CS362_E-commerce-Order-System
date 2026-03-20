package response

// SuccessResponse wraps a successful API response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorDetail holds structured error information
type ErrorDetail struct {
	Code       string `json:"code"`
	Message    string `json:"message,omitempty"`
	ProductID  string `json:"productId,omitempty"`
	RetryAfter int    `json:"retryAfter,omitempty"`
}

// ErrorResponse wraps a failed API response
type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   ErrorDetail `json:"error"`
}

// NewSuccess returns a 200/201 response envelope
func NewSuccess(data interface{}) *SuccessResponse {
	return &SuccessResponse{Success: true, Data: data}
}

// NewError returns an error response envelope
func NewError(code, message string) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Error:   ErrorDetail{Code: code, Message: message},
	}
}
