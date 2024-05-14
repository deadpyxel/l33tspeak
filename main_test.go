package main

import (
	"reflect"
	"testing"
)

func TestLeetSpeak(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		ignoreList string
		expected   string
		replPerc   float64
		randSeed   int64
	}{
		{name: "Basic input", input: "hello world", ignoreList: "", expected: "h3ll0 w0Rld", replPerc: 1.0, randSeed: 1},
		{name: "Partial replacement", input: "hello world", ignoreList: "", expected: "heLL0 W0RlD", replPerc: 0.5, randSeed: 1},
		{name: "When replacement percentage is zero no swaps are made from letter to number", input: "hello world", ignoreList: "", expected: "heLLO wOrLD", replPerc: 0.0, randSeed: 1},
		{name: "When passing string with special characters they are kept", input: "-helLo w0rld!", ignoreList: "", expected: "-h3lL0 w0rLD!", replPerc: 1.0, randSeed: 1},
		{name: "When passing an empty string an empty string is returned", input: "", ignoreList: "", expected: "", replPerc: 1.0, randSeed: 1},
		{name: "When passing an string with an ignore list ignored characters are returned as is", input: "hello", ignoreList: "h,e,l,o", expected: "hello", replPerc: 1.0, randSeed: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeetSpeak(tt.input, tt.ignoreList, tt.replPerc, tt.randSeed); got != tt.expected {
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
			if got := UnleetSpeak(tt.input); got != tt.expected {
				t.Errorf("unleetSpeak() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCreateIgnoreList(t *testing.T) {
	t.Run("When receiving properly formatted comma separated input results in a slice of strings", func(t *testing.T) {
		tests := []struct {
			input    string
			expected []string
		}{
			{"a,b,c", []string{"a", "b", "c"}},
			{"1,2,3", []string{"1", "2", "3"}},
			{"x,y,z", []string{"x", "y", "z"}},
		}

		for _, test := range tests {
			result := createIgnoreList(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		}
	})
}

// Benchmark tests
func BenchmarkLeetSpeak(b *testing.B) {
	b.Run("Simple case without ignores", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LeetSpeak("hello world", "", 1.0, 1)
		}
	})
	b.Run("Simple case with ignores", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LeetSpeak("hello world", "h,e,l,o,w,r,d", 1.0, 1)
		}
	})
}

func BenchmarkUnleetSpeak(b *testing.B) {
	b.Run("Simple case", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			UnleetSpeak("h3ll0 w0rld")
		}
	})
}
