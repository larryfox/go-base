package base

import (
	"fmt"
)

type err struct {
	string
}

func (e err) Error() string {
	return e.string
}

var (
	NegativeIntegerError = err{"Integer can't be negative"}
	EmptyStringError     = err{"String can't be empty"}
	InvalidBaseError     = newErrorf("Invalid base %d. Base must be greater than 1 and less than 63")
)

func newErrorf(format string) func(int) error {
	return func(n int) error {
		return err{fmt.Sprintf(format, n)}
	}
}
