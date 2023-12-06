package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{2, 2, 4},
		{3, 3, 6},
		{3, 5, 8},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Add(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestSubstract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{4, 2, 2},
		{5, 3, 2},
		{6, 4, 2},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Subtract(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{2, 2, 4},
		{3, 3, 9},
		{3, 5, 15},
	}

	for _, tc := range testCases {
		got  := calculator.Multiply(tc.a, tc.b)

		if got != tc.want {
			t.Errorf("Multiply(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}
