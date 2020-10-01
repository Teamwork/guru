package guru

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func stackTrace() errors.StackTrace {
	var pcs [8]uintptr
	n := runtime.Callers(1, pcs[:])
	var st stack = pcs[0:n]
	return st.StackTrace()
}

func TestStackTraceFormat(t *testing.T) {
	tests := []struct {
		errors.StackTrace
		format string
		want   string
	}{
		{
			stackTrace()[:2],
			"%s",
			`[stack_test.go stack_test.go]`,
		},
		{
			stackTrace()[:2],
			"%v",
			`[stack_test.go:15 stack_test.go:32]`,
		},
		{
			stackTrace()[:2],
			"%+v",
			"\n" +
				"github.com/teamwork/guru.stackTrace\n" +
				"\tgithub.com/teamwork/guru/stack_test.go:15\n" +
				"github.com/teamwork/guru.TestStackTraceFormat\n" +
				"\tgithub.com/teamwork/guru/stack_test.go:37",
		},
		{
			stackTrace()[:2],
			"%#v",
			`[]errors.Frame{stack_test.go:15, stack_test.go:46}`,
		},
	}

	cwd, _ := os.Getwd()
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := fmt.Sprintf(tt.format, tt.StackTrace)
			got = strings.ReplaceAll(got, cwd, "github.com/teamwork/guru")
			if got != tt.want {
				t.Errorf("\ngot:  %q\nwant: %q", got, tt.want)
			}
		})
	}
}
