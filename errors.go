package errors

import (
	"fmt"
)

//---------------------------------------------------------------------------------------------
func NewStackTraceableError(message string, cause error) error {
	return &stackTraceableError{
		msg:   message,
		cause: cause,
		stack: callers(),
	}
}

func NewStackTraceableErrorF(format string, args ...interface{}) error {
	return &stackTraceableError{
		msg:   fmt.Sprintf(format, args...),
		stack: callers(),
	}
}


// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return err
}
