//go:build go1.20

package errors

import "errors"

// Join returns an error that wraps the given errors.
// Any nil error values are discarded.
// Join returns nil if errs contains no non-nil values.
// The error formats as the concatenation of the strings obtained
// by calling the Error method of each element of errs, with a newline
// between each string.
func Join(errs ...error) error {
	err := errors.Join(errs...)
	if err == nil {
		return nil
	}
	return newErr(err, "")
}
