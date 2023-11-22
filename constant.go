package errors

type Constant string

func (err Constant) Error() string {
	return string(err)
}

func (err Constant) Is(other error) bool {
	return other.Error() == err.Error()
}
