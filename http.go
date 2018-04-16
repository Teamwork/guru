package guru

// HTTPUserError reports if this HTTP status code is a user error (i.e. in the
// 4xx range).
func HTTPUserError(err error) bool {
	// TODO: decide what to do with this; it's useful enough, but also a
	// remanent from the httperr days.
	if err == nil {
		return false
	}

	code := Code(err)
	return code >= 400 && code <= 499
}
