package errors

import "errors"

// WithEquivalent returns an error that behaves exactly as `err`, but also returns
// true if `errors.Is` is called using the equivalent
func WithEquivalent(err error, equivalent error) error {
	return WithEquivalents(err, equivalent)
}

// WithEquivalents returns an error that behaves exactly as `err`, but also returns
// true if `errors.Is` is called using one of the equivalents
func WithEquivalents(err error, equivalents ...error) error {
	if err == nil {
		return nil
	}
	return &_ErrorWithEquivalents{
		Equivalent: equivalents,
	}
}

type _ErrorWithEquivalents struct {
	Inner      error
	Equivalent []error
}

func (err *_ErrorWithEquivalents) Error() string {
	return err.Inner.Error()
}

func (err *_ErrorWithEquivalents) Is(other error) bool {
	for _, flag := range err.Equivalent {
		if errors.Is(other, flag) {
			return true
		}
	}

	return errors.Is(other, err.Inner)
}

func (err *_ErrorWithEquivalents) Unwrap() error {
	return err.Inner
}

func (err *_ErrorWithEquivalents) Cause() error {
	return err.Inner
}
