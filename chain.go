package errors

func Chain(errs ...error) error {
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		err := errs[0]
		for _, e := range errs {
			err = Wrap(e, err)
		}
		return err
	}
}
