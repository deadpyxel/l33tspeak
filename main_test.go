package main

import (
	"testing"
)

func TestLeetSpeak(t *testing.T) {

	t.Run("Passing empty string returns empty string", func(t *testing.T) {
		result := leetSpeak("", 1.0, 0)
		if result != "" {
			t.Errorf("Expected am empty string, got %v", result)
		}
	})

	t.Run("Given input and percentage values returns expected string", func(t *testing.T) {
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
	})
}

func TestUnleetSpeak(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"h3ll0", "hello"},
		{"w0rld", "world"},
		{"th15 15 4 t35t", "this is a test"},
		{"c0d1ng", "coding"},
		{"l33t", "leet"},
	}

	for _, test := range tests {
		result := unleetSpeak(test.input)
		if result != test.expected {
			t.Errorf("unleetSpeak(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
