package errors

import (
	"errors"
	"fmt"
)

type _WrappedError struct {
	Inner   error
	Wrapper error
}

func Wrap(err error, wrapper error) error {
	return &_WrappedError{
		Inner:   err,
		Wrapper: wrapper,
	}
}

func (err *_WrappedError) Error() string {
	return fmt.Sprintf("%s: %s", err.Wrapper.Error(), err.Inner.Error())
}

func (err *_WrappedError) Is(other error) bool {
	if errors.Is(other, err.Wrapper) {
		return true
	}
	return errors.Is(other, err.Inner)
}

func (err *_WrappedError) Unwrap() error {
	return err.Inner
}

func (err *_WrappedError) Cause() error {
	return err.Inner
}
