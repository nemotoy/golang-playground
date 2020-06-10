package error

type ExError struct {
	Type string
	Err  error
}

func (e *ExError) Error() string { return e.Err.Error() }

func (e *ExError) UnWrap() error { return e.Err }

func IsExErr(err error) bool {
	_, exist := err.(*ExError)
	return exist
}

func ClassifyErr(err error) error {
	switch e := err.(type) {
	case *ExError:
		return e.Err
	}
	return err
}
