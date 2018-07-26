package gateerror

import "fmt"

// Error for gateway
type Error struct {
	code      int
	message   string
	interCode int
}

// New Error instance
func New(code int, message string, interCode int) *Error {
	return &Error{code, message, interCode}
}

// Error getter
func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s (%d)", e.code, e.message, e.interCode)
}

// Code getter
func (e *Error) Code() int {
	return e.code
}

// Message getter
func (e *Error) Message() string {
	return e.message
}

// InterCode getter
func (e *Error) InterCode() int {
	return e.interCode
}
