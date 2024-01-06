package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO WORLD",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range testCases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Lengths are different. %v is expected, but received %v",
				len(cs.expected),
				len(actual),
			)
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]

			if actualWord != expectedWord {
				t.Errorf("Words are not equal, %v is not equal to %v",
					actualWord,
					expectedWord,
				)
			}
		}
	}

}
