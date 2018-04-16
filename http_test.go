package guru

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestHTTPUserError(t *testing.T) {
	tests := []struct {
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

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := HTTPUserError(tt.in)
			if out != tt.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tt.want)
			}
		})
	}
}
