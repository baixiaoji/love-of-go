package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func closeEnough(a, b float64, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
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
		got := calculator.Multiply(tc.a, tc.b)

		if got != tc.want {
			t.Errorf("Multiply(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{1, 1, 1},
		{-1, -1, 1},
		{10, 2, 5},
		{1, 3, 0.33333333333333},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for the valid input %v", err)
		}
		if !closeEnough(got, tc.want, 1e-9) {
			t.Errorf("Divide(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}

}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)

	if err == nil {
		t.Fatalf("want error for the invalid input")
	}
}
