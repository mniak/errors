package errors

import (
	"errors"
)

// New creates a new error with the specified message
func New(msg string) error {
	return errors.New(msg)
}

// func Newf(format string, args ...any) error {
// 	return fmt.Errorf(format, args...)
// }
