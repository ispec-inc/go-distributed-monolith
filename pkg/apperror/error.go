package apperror

import (
	"github.com/pkg/errors"
)

type Error struct {
	code Code
	msg  string
	org  error
}

func New(code Code, msg string) error {
	return errors.WithStack(&Error{
		code: code,
		msg:  msg,
	})
}

func Wrap(code Code, err error) error {
	return errors.WithStack(&Error{
		code: code,
		msg:  err.Error(),
		org:  err,
	})
}

func (e *Error) Code() Code {
	return e.code
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Origin() error {
	return e.org
}

func Unwrap(err error) *Error {
	if err == nil {
		return nil
	}

	aerr, ok := errors.Cause(err).(*Error)
	if ok {
		aerr.msg = err.Error()
		return aerr
	}
	return nil
}
