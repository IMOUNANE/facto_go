package factgorial

import (
	"errors"
	"gitlab.eemi.tech/golang/factgorial/messages"
	 stderrors "errors"
)


func Factorial(n int) (uint64, error) {
	if n < 0 {
		return 0, errors.New(messages.Negative)
	}

	if n < 2 {
		return 1, nil
	}

	/// unsigned n
	un := uint64(n)
	m := uint64(1)

	for ; un > 1; un-- {
		m = m * un
		if m == 0 {
			return m, errors.New(messages.OverflowError)
		}
	}

	return m, nil
}

Factorial(10)
