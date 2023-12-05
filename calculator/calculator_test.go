package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if got != want {
		t.Errorf("Add(2, 2) = %v, want %v", got, want)
	}
}

func TestSubstract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if got != want {
		t.Errorf("Substract(4, 2) = %v, want %v", got, want)
	}
}
