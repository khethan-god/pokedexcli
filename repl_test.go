package main

import (
	"testing" // Import the testing package
)

// unit test should focus on testing individual functions or small units of code in isolation

func TestCleanInput(t *testing.T) {

	// arrage/set up the necessary inputs and pre conditions for testing
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "gOlang",
			expected: []string{"golang"},
		},
		{
			input:    "Multiple   spaces   here",
			expected: []string{"multiple", "spaces", "here"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " single ",
			expected: []string{"single"},
		},
	}

	for _, c := range cases {
		// call the funciton being tested (ACT)
		actual := cleanInput(c.input)
		// assert the conditions and result (Assert)
		if len(actual) != len(c.expected) {
			t.Errorf("Test Case Failed: Input: '%s', Expected Length: %d, Actual Length: %d",
				c.input, len(c.expected), len(actual))
			continue // Move to the next test case
		}
		for i := range actual {
			// assert the conditions and result (Assert)
			if actual[i] != c.expected[i] {
				t.Errorf("Test Case Failed: Input: '%s', Expected Word at Index %d: '%s', Actual Word: '%s'",
					c.input, i, c.expected[i], actual[i])
			}
		}
	}
}

/*
In Go, functions and variables declared within the same package are accessible to all other
files within that same directory without needing to be explicitly exported.
*/
