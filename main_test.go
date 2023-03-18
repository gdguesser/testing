package main

import (
	"testing"
)

func TestSomeValue(t *testing.T) {
	result := SomeValue()
	if result != "expected value" {
		t.Errorf("SomeValue() = %q, want %q", result, "expected value")
	}
}

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, esperado int
	}{
		{2, 3, 6},
		{-2, 3, -6},
		{2, -3, -6},
		{0, 3, 0},
	}

	for _, test := range tests {
		result := Multiply(test.a, test.b)
		if result != test.esperado {
			t.Errorf("Multiply(%d, %d) = %d; want %d", test.a, test.b, result, test.esperado)
		}
	}
}
