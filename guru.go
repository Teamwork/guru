// Package guru provides Go errors with a Guru Meditation Code.
package guru // import "github.com/teamwork/guru"

import (
	"fmt"

	"github.com/pkg/errors"
)

// coder is the main interface to errors in this package.
type coder interface {
	Code() int
}

// causer is a local version of github.com/pkg/errors.causer
type causer interface {
	Cause() error
}

// stackTracer is a local version of github.com/pkg/errors.stackTracer
type stackTracer interface {
	StackTrace() errors.StackTrace
}

type withCode struct {
	error
	code int
	*stack
}

func (e *withCode) Cause() error                 { return e.error }
func (e *withCode) Code() int                    { return e.code }
func (e withCode) Format(s fmt.State, verb rune) { fmt.Fprintf(s, "error %v: %v", e.code, e.error) } // nolint: errcheck

type wrapped struct {
	msg  string
	code int
	error
	*stack
}

func (e *wrapped) Error() string { return e.msg }
func (e *wrapped) Cause() error  { return e.error }
func (e *wrapped) Code() int     { return e.code }
func (e wrapped) Format(s fmt.State, verb rune) {
	fmt.Fprintf(s, "error %v: %v", e.code, e.error) // nolint: errcheck
	if e.msg != "" {
		fmt.Fprintf(s, ": %v", e.msg) // nolint: errcheck
	}
}

// New returns a new error message with a stack trace and error code.
func New(code int, msg string) error {
	return &withCode{
		error: errors.New(msg),
		code:  code,
		stack: callers(),
	}
}

// Errorf returns a new error message with stack trace and error code.
func Errorf(code int, format string, args ...interface{}) error {
	return &withCode{
		error: fmt.Errorf(format, args...),
		code:  code,
		stack: callers(),
	}
}

// WithCode wraps an existing error with the provided error code and stack
// trace. It will return nil if err is nil.
func WithCode(code int, err error) error {
	if err == nil {
		return nil
	}
	return &withCode{
		error: err,
		code:  code,
		stack: callers(),
	}
}

// Wrap returns an error annotating err with a stack trace, error code, and the
// supplied message. It will return nil if err is nil.
func Wrap(code int, err error, msg string) error {
	if err == nil {
		return nil
	}
	return &wrapped{
		msg:   msg,
		code:  code,
		error: err,
		stack: callers(),
	}
}

// Wrapf returns an error annotating err with a stack trace, error code, and the
// format specifier. It will return nil if err is nil.
func Wrapf(code int, err error, msg string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &wrapped{
		msg:   fmt.Sprintf(msg, args...),
		code:  code,
		error: err,
		stack: callers(),
	}
}

// Code extracts the highest-level error code from the error or the errors it
// wraps. It will return 0 if the error does not implement the coder interface.
func Code(err error) int {
	for {
		if sc, ok := err.(coder); ok {
			return sc.Code()
		}
		if cause, ok := err.(causer); ok {
			err = cause.Cause()
		} else {
			break
		}
	}
	return 0
}
