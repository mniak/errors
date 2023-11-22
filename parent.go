package errors

import "fmt"

type _ErrorWithParent struct {
	Parent error
	Child  error
}

func WithParent(err error, parentError error) error {
	return _ErrorWithParent{
		Parent: parentError,
		Child:  err,
	}
}

func (err _ErrorWithParent) Error() string {
	return fmt.Sprintf("%s: %s", err.Parent.Error(), err.Child.Error())
}
