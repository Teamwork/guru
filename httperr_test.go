package httperr

import (
	"errors"
	"fmt"
	"testing"

	"github.com/labstack/echo"
)

func TestFormatWithStatus(t *testing.T) {
	cases := []struct {
		in   withStatus
		fmt  string
		want string
	}{
		{
			withStatus{
				error:  errors.New("oh noes"),
				status: 42,
			},
			"%v",
			"oh noes",
		},
		{
			withStatus{
				error:  errors.New("oh noes"),
				status: 42,
			},
			"%#v",
			"oh noes",
		},
		{
			withStatus{
				error:  errors.New("oh noes"),
				status: 42,
			},
			"%s",
			"oh noes",
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

func TestFormatWrapped(t *testing.T) {
	cases := []struct {
		in   wrapped
		fmt  string
		want string
	}{
		{
			wrapped{
				error:  errors.New("oh noes"),
				status: 42,
			},
			"%v",
			"oh noes",
		},
		{
			wrapped{
				error:  errors.New("oh noes"),
				status: 42,
			},
			"%#v",
			"oh noes",
		},
		{
			wrapped{
				error:  errors.New("oh noes"),
				status: 42,
			},
			"%s",
			"oh noes",
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

func TestStatusCode(t *testing.T) {
	cases := []struct {
		in   error
		want int
	}{
		{errors.New("foo"), 0},
		{New(42, "foo"), 42},
		{echo.NewHTTPError(400, ""), 400},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := StatusCode(tc.in)
			if out != tc.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tc.want)
			}
		})
	}
}

func TestUserError(t *testing.T) {
	cases := []struct {
		in   error
		want bool
	}{
		{nil, false},
		{errors.New("asd"), false},
		{New(500, "asd"), false},
		{New(399, "asd"), false},
		{New(0, "asd"), false},

		{New(400, "asd"), true},
		{New(499, "asd"), true},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := UserError(tc.in)
			if out != tc.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tc.want)
			}
		})
	}
}
