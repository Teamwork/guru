package httperr

import (
	"fmt"

	"github.com/pkg/errors"
)

// statusCoder is the main interface to errors in this package.
type statusCoder interface {
	StatusCode() int
}

// causer is a local version of github.com/pkg/errors.causer
type causer interface {
	Cause() error
}

// stackTracer is a local version of github.com/pkg/errors.stackTracer
type stackTracer interface {
	StackTrace() errors.StackTrace
}

// withStatus wraps a standard error and an HTTP status code.
type withStatus struct {
	error
	status int
	*stack
}

var _ error = &withStatus{}
var _ statusCoder = &withStatus{}
var _ causer = &withStatus{}
var _ stackTracer = &withStatus{}

func (e *withStatus) Cause() error {
	return e.error
}

func (e *withStatus) StatusCode() int {
	return e.status
}

// New returns a new error message, tied to the provided HTTP status code.
func New(status int, msg string) error {
	return &withStatus{
		error:  errors.New(msg),
		status: status,
		stack:  callers(),
	}
}

// Errorf returns a new error message, tied to the provided HTTP status code.
func Errorf(status int, format string, args ...interface{}) error {
	return &withStatus{
		error:  fmt.Errorf(format, args...),
		status: status,
		stack:  callers(),
	}
}

// WithStatus wraps an existing error with the provided HTTP status code.
func WithStatus(status int, err error) error {
	return &withStatus{
		error:  err,
		status: status,
		stack:  callers(),
	}
}
