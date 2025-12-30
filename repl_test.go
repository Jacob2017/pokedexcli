package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "howdy gang",
			expected: []string{"howdy", "gang"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// fmt.Println(actual)
		// fmt.Println(c.expected)
		if len(actual) != len(c.expected) {
			t.Errorf("Slice received has different length (%d) to expected (%d)", len(actual), len(c.expected))
		} else {
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("Word %d (%s) did not match expected (%s)", i, word, expectedWord)
				}
			}
		}

	}
}
