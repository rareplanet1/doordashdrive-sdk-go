package doordashdrive

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode  int            `json:"status_code"`
	Code        string         `json:"code,omitempty"`
	Message     string         `json:"message,omitempty"`
	FieldErrors []FieldError   `json:"field_errors,omitempty"` // field validation errors, if any
	Response    *http.Response `json:"response,omitempty"`     // the full response that produced the error
}

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error: status=%d, code=%s, message=%s", e.StatusCode, e.Code, e.Message)
}
