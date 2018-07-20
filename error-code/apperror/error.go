package apperror

import "fmt"

// error codes
const (
	_ = iota
	ErrBadRequestData
	ErrNotFound
)

// AppError struct with code and message
type AppError struct {
	code    int
	message string
	err     error
}

// New create Error instance
func New(code int, message string, err error) *AppError {
	return &AppError{code, message, err}
}

// Error getter
func (e *AppError) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.message)
}

// Code getter
func (e *AppError) Code() int {
	return e.code
}

// Message getter
func (e *AppError) Message() string {
	return e.message
}

// ErrorSummry getter
func (e *AppError) ErrorSummry() string {
	return fmt.Sprintf("%d: %s, [%s]", e.code, e.message, e.err.Error())
}
