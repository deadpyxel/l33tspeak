package main

import (
	"testing"
)

func TestLeetSpeak(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		replPerc float64
		randSeed int64
	}{
		{name: "Basic input", input: "hello world", expected: "h3ll0 w0Rld", replPerc: 1.0, randSeed: 1},
		{name: "Partial replacement", input: "hello world", expected: "heLL0 W0RlD", replPerc: 0.5, randSeed: 1},
		{name: "When replacement percentage is zero no swaps are made from letter to number", input: "hello world", expected: "heLLO wOrLD", replPerc: 0.0, randSeed: 1},
		{name: "When passing string with special characters they are kept", input: "-helLo w0rld!", expected: "-h3lL0 w0rLD!", replPerc: 1.0, randSeed: 1},
		{name: "When passing an empty string an empty string is returned", input: "", expected: "", replPerc: 1.0, randSeed: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leetSpeak(tt.input, tt.replPerc, tt.randSeed); got != tt.expected {
				t.Errorf("Expected [%v] for leetSpeak(\"%v\", %v, %v), got %v instead.", tt.expected, tt.input, tt.replPerc, tt.randSeed, got)
			}
		})
	}
}

func TestUnleetSpeak(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Basic Test", "h3ll0 w0rld", "hello world"},
		{"When input has special characters they are kept as is", "G0 15 fUn!", "Go is fUn!"},
		{"When an input has no replacements to be made it is returned as is", "hello world", "hello world"},
		{"When an input is am empty string an empty string is returned", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unleetSpeak(tt.input); got != tt.expected {
				t.Errorf("unleetSpeak() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkLeetSpeak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetSpeak("hello world", 1.0, 1)
	}
}

func BenchmarkUnleetSpeak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unleetSpeak("h3ll0 w0rld")
	}
}
