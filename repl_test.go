package main

import "testing"

// syntax for single test
func TestCleanInput(t *testing.T) {
	len_actual := len(cleanInput("Hello"))
	if len_actual != 1 {
		t.Errorf("Lengths are not the same %v vs %v", len_actual, 1)
	}
}

// syntax for table-driven test
func TestCleanInputTableDriven(t *testing.T) {
	var tests = []struct {
		input    string
		expected []string
	}{
		{"Hello", []string{"hello"}},
		{"Hello World", []string{"hello", "world"}},
	}

	for _, test := range tests {
		actual := cleanInput(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("Output is not the same. Expected len: %v, Actual len: %v", len(test.expected), len(actual))
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := test.expected[i]
			if actualWord != expectedWord {
				t.Errorf("Words don't match. Expected: %v vs Actual: %v", expectedWord, actualWord)
			}
		}
	}
}
