package httperr

import (
	"errors"
	"fmt"
	"testing"
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
			"oh noes (42)",
		},
		{
			withStatus{
				error:  errors.New("oh noes"),
				status: 42,
			},
			"%#v",
			"oh noes (42)",
		},
		{
			withStatus{
				error:  errors.New("oh noes"),
				status: 42,
			},
			"%s",
			"oh noes (42)",
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
			"oh noes (42)",
		},
		{
			wrapped{
				error:  errors.New("oh noes"),
				status: 42,
			},
			"%#v",
			"oh noes (42)",
		},
		{
			wrapped{
				error:  errors.New("oh noes"),
				status: 42,
			},
			"%s",
			"oh noes (42)",
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
