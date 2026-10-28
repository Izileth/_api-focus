package error

import (
	"fmt"
	"net/http"
)

// AppError represents a structured application error
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError with the specified HTTP status code and message
func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Wrap wraps an existing error into an AppError
func Wrap(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// BadRequest creates an AppError for 400 Bad Request
func BadRequest(message string) *AppError {
	return New(http.StatusBadRequest, message)
}

// NotFound creates an AppError for 404 Not Found
func NotFound(message string) *AppError {
	return New(http.StatusNotFound, message)
}

// Unauthorized creates an AppError for 401 Unauthorized
func Unauthorized(message string) *AppError {
	return New(http.StatusUnauthorized, message)
}

// Internal creates an AppError for 500 Internal Server Error
func Internal(message string, err error) *AppError {
	return Wrap(http.StatusInternalServerError, message, err)
}
