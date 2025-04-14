package errors

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestLineAndFormatting(t *testing.T) {
	err := New("an error")
	wrapErr := Wrap(err, "another error")
	assert.Equal(t, `an error`, fmt.Sprintf("%s", err))
	assert.Equal(t, `"an error"`, fmt.Sprintf("%q", err))
	assert.Equal(t, `errors_test.go:11: an error`, fmt.Sprintf("%+v", err))
	assert.Equal(t, `another error: an error`, fmt.Sprintf("%s", wrapErr))
	assert.Equal(t, `errors_test.go:12: another error: errors_test.go:11: an error`, fmt.Sprintf("%+v", wrapErr))
}

func TestUnwrapAllAndInnermost(t *testing.T) {
	err := Wrap(Join(New("A"), Wrap(New("B"), "C")), "D")
	errs := UnwrapAll(err)
	errstrings := make([]string, len(errs))
	innermost := make([]bool, len(errs))
	for i, err := range errs {
		errstrings[i] = err.Error()
		innermost[i] = Innermost(err)
	}
	assert.Equal(t, []string{"A", "B", "C: B", "D: A\nC: B"}, errstrings)
	assert.Equal(t, []bool{true, true, false, false}, innermost)
}
