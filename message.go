package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

func WithMessage(err error, message string) error {
	return errors.WithMessage(err, message)
}

func WithMessagef(err error, format string, args ...any) error {
	return errors.WithMessagef(err, format, args...)
}

// WithCustomMessage replaces the error message but maintains other parts of the behavior, such as equality
func WithCustomMessage(err error, customMessage string) error {
	return &_ErrorWithCustomMessage{
		Inner:         err,
		CustomMessage: customMessage,
	}
}

// WithCustomMessagef replaces the error message but maintains other parts of the behavior, such as equality
func WithCustomMessagef(err error, customMessageFormat string, args ...any) error {
	return &_ErrorWithCustomMessage{
		Inner:         err,
		CustomMessage: fmt.Sprintf(customMessageFormat, args...),
	}
}

type _ErrorWithCustomMessage struct {
	Inner         error
	CustomMessage string
}

func (e *_ErrorWithCustomMessage) Error() string {
	return e.CustomMessage
}

func (e *_ErrorWithCustomMessage) Is(other error) bool {
	return errors.Is(e.Inner, other)
}
