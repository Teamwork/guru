package guru

import (
	"errors"
	"fmt"
	"testing"
)

var _ error = &withCode{}
var _ coder = &withCode{}
var _ causer = &withCode{}
var _ stackTracer = &withCode{}

var _ error = &wrapped{}
var _ stackTracer = &wrapped{}
var _ coder = &wrapped{}
var _ causer = &wrapped{}

func TestFormatWithStatus(t *testing.T) {
	tests := []struct {
		in   withCode
		fmt  string
		want string
	}{
		{
			withCode{
				error: errors.New("oh noes"),
				code:  42,
			},
			"%v",
			"error 42: oh noes",
		},
		{
			withCode{
				error: errors.New("oh noes"),
				code:  42,
			},
			"%#v",
			"error 42: oh noes",
		},
		{
			withCode{
				error: errors.New("oh noes"),
				code:  42,
			},
			"%s",
			"error 42: oh noes",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := fmt.Sprintf(tt.fmt, tt.in)
			if out != tt.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tt.want)
			}
		})
	}
}

func TestFormatWrapped(t *testing.T) {
	tests := []struct {
		in   wrapped
		fmt  string
		want string
	}{
		{
			wrapped{
				error: errors.New("oh noes"),
				code:  42,
			},
			"%v",
			"error 42: oh noes",
		},
		{
			wrapped{
				error: errors.New("oh noes"),
				code:  42,
			},
			"%#v",
			"error 42: oh noes",
		},
		{
			wrapped{
				error: errors.New("oh noes"),
				code:  42,
			},
			"%s",
			"error 42: oh noes",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := fmt.Sprintf(tt.fmt, tt.in)
			if out != tt.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tt.want)
			}
		})
	}
}

func TestStatusCode(t *testing.T) {
	tests := []struct {
		in   error
		want int
	}{
		{errors.New("foo"), 0},
		{New(42, "foo"), 42},
		{Wrap(666, New(42, "foo"), "bar"), 666},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := Code(tt.in)
			if out != tt.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tt.want)
			}
		})
	}
}
