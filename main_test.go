package main

import (
	"reflect"
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

func TestGetCEPInfo(t *testing.T) {
	// Define a slice of test cases
	tests := []struct {
		cep        string
		expected   *CEPInfo
		expectFail bool
	}{
		{
			// First test case
			cep: "01001000",
			expected: &CEPInfo{
				Logradouro:  "Praça da Sé",
				Complemento: "lado ímpar",
				Bairro:      "Sé",
				Localidade:  "São Paulo",
				UF:          "SP",
			},
			expectFail: false, // We expect this test case to succeed
		},
		{
			// Second test case
			cep:        "invalid_cep",
			expected:   nil,
			expectFail: true, // We expect this test case to fail
		},
	}

	// Iterate over the test cases using a for loop
	for _, tc := range tests {
		// Use t.Run to create a sub-test for each test case
		t.Run(tc.cep, func(t *testing.T) {
			// Call the function being tested with the current test case's cep field
			result, err := GetCEPInfo(tc.cep)

			// If expectFail is true, we expect an error to be returned
			if tc.expectFail {
				if err == nil {
					t.Fatalf("expected an error but got nil")
				}
				return
			}

			// If we expect the function to succeed, but it returns an error, fail the test
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			// Compare the function's return value to the expected value using reflect.DeepEqual
			if !reflect.DeepEqual(result, tc.expected) {
				t.Fatalf("expected %+v, but got %+v", tc.expected, result)
			}
		})
	}
}
