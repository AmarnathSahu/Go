package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(5, 6)
	expected := 11

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestAddInBulk(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{2, 4, 6},
		{5, 5, 10},
	}

	for _, test := range tests {
		result := test.a + test.b

		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}
