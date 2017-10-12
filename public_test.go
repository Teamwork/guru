package httperr

import (
	"errors"
	"fmt"
	"testing"
)

func TestFormatPublicError(t *testing.T) {
	cases := []struct {
		in   publicError
		fmt  string
		want string
	}{
		{
			publicError{
				err: errors.New("oh noes"),
				msg: "public",
			},
			"%v",
			"public",
		},
		{
			publicError{
				err: errors.New("oh noes"),
				msg: "public",
			},
			"%#v",
			"public",
		},
		{
			publicError{
				err: errors.New("oh noes"),
				msg: "public",
			},
			"%s",
			"public",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := fmt.Sprintf(tc.fmt, tc.in)
			if out != tc.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tc.want)
			}
		})
	}
}

func TestFormatPublicWithStatus(t *testing.T) {
	cases := []struct {
		in   publicWithStatus
		fmt  string
		want string
	}{
		{
			publicWithStatus{
				status: 42,
				publicError: publicError{
					err: errors.New("oh noes"),
					msg: "public",
				},
			},
			"%v",
			"public",
		},
		{
			publicWithStatus{
				status: 42,
				publicError: publicError{
					err: errors.New("oh noes"),
					msg: "public",
				},
			},
			"%#v",
			"public",
		},
		{
			publicWithStatus{
				status: 42,
				publicError: publicError{
					err: errors.New("oh noes"),
					msg: "public",
				},
			},
			"%s",
			"public",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := fmt.Sprintf(tc.fmt, tc.in)
			if out != tc.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tc.want)
			}
		})
	}
}
