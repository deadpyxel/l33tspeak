package main

import (
	"testing"
)

func TestLeetSpeak(t *testing.T) {

	tests := []struct {
		input    string
		replPerc float64
		expected string
	}{
		{"hello", 0.5, "hElL0"},
		{"world", 0.8, "w0rlD"},
		{"leet", 0.2, "lE3t"},
		{"testing", 0.0, "tEsTinG"},
	}

	for _, test := range tests {
		result := leetSpeak(test.input, test.replPerc, 42)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s for input %s and replPerc %f", test.expected, result, test.input, test.replPerc)
		}
	}
}
