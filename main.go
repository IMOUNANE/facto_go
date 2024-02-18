package main

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

var (
	ErrNegativeInput = errors.New("negative input")
	ErrOverflow      = errors.New("factorial overflow")
)

func Factorial(n int) (uint64, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	if n > 20 {
		return 0, ErrOverflow
	}

	result := uint64(1)
	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}

	return result, nil
}

func RecursiveFactorial(n int) (uint64, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}
	if n > 20 {
		return 0, ErrOverflow
	}

	if n == 0 {
		return 1, nil
	}

	fact, _ := RecursiveFactorial(n - 1)
	if fact > math.MaxUint64/uint64(n) {
		return 0, ErrOverflow
	}

	return fact * uint64(n), nil
}

func BenchmarkFactorial(b *testing.B, f func(int) (uint64, error)) {
	for i := 0; i < b.N; i++ {
		f(10) // Benchmark with input 10
	}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input    int
		expected uint64
		err      error
	}{
		{0, 1, nil},
		{1, 1, nil},
		{5, 120, nil},
		{-1, 0, ErrNegativeInput},
		{21, 0, ErrOverflow},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input_%d", tc.input), func(t *testing.T) {
			result, err := Factorial(tc.input)
			if err != tc.err {
				t.Errorf("Expected error: %v, got: %v", tc.err, err)
			}
			if result != tc.expected {
				t.Errorf("Expected result: %d, got: %d", tc.expected, result)
			}
		})
	}
}

func TestRecursiveFactorial(t *testing.T) {
	testCases := []struct {
		input    int
		expected uint64
		err      error
	}{
		{0, 1, nil},
		{1, 1, nil},
		{5, 120, nil},
		{-1, 0, ErrNegativeInput},
		{21, 0, ErrOverflow},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input_%d", tc.input), func(t *testing.T) {
			result, err := RecursiveFactorial(tc.input)
			if err != tc.err {
				t.Errorf("Expected error: %v, got: %v", tc.err, err)
			}
			if result != tc.expected {
				t.Errorf("Expected result: %d, got: %d", tc.expected, result)
			}
		})
	}
}

func main() {
	// Run tests
	tests := []testing.InternalTest{
		{"TestFactorial", TestFactorial},
		{"TestRecursiveFactorial", TestRecursiveFactorial},
	}
	testing.Main(nil, tests, nil, nil)
}
