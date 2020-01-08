[![Build Status](https://travis-ci.com/Teamwork/guru.svg?branch=master)](https://travis-ci.com/Teamwork/guru)
[![Codecov](https://img.shields.io/codecov/c/github/Teamwork/guru.svg?style=flat)](https://codecov.io/gh/Teamwork/guru)
[![GoDoc](https://godoc.org/github.com/Teamwork/guru?status.svg)](http://godoc.org/github.com/Teamwork/guru)

The guru package allows adding a Guru Meditation Number to errors:

```go
// Error constants.
const (
	CodeFruitOverflow = iota + 1
	CodeBoozeUnderrun
	CodeExpired
)

func Example() {
	// Construct a new error.
	err := guru.New(CodeFruitOverflow, "too many bananas")
	fmt.Println(err) // error 1: too many bananas

	// Retrieve the error code
	code := guru.Code(err)
	fmt.Println(code) // 1

	// Add error code to existing error.
	err = errors.New("not enough beer")
	err = guru.WithCode(CodeBoozeUnderrun, err)
	fmt.Println(err) // error 2: not enough beer

	// Add error code to existing error with context.
	err = errors.New("Dennis Ritchie")
	err = guru.Wrap(CodeExpired, err, "no longer with us")
	fmt.Println(err) // error 3: Dennis Ritchie: no longer with us

	// For HTTP applications, it may be useful to directly the HTTP status codes:
	err = guru.New(http.StatusNotAcceptable, "Justin Bieber")
	fmt.Println(err) // error 406: Justin Bieber

	// Error codes can be overriden:
	err = guru.New(1, "oh noes")
	err = guru.WithCode(2, err)
	fmt.Println(guru.Code(err)) // 2
}
```

guru is built on top of [github.com/pkg/errors](https://github.com/pkg/errors);
all errors that guru returns are github.com/pkg/errors.
