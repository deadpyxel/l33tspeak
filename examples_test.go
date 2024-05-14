package main

import (
	"fmt"
	"testing"
)

func ExampleLeetSpeak() {
	input := "hello world"
	ignoreList := "o"
	replPerc := 0.5
	randSeed := int64(12345)
	result := LeetSpeak(input, ignoreList, replPerc, randSeed)
	fmt.Println(result)
	// Output: hElLo woRLd
}

func TestLeetSpeakExampleUsage(t *testing.T) {
	// This test runs the example function as a test.
	ExampleLeetSpeak()
}
