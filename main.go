package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func leetSpeak(input string, replPerc float64, randSeed int64) string {
	replacements := map[string]string{
		"a": "4",
		"e": "3",
		"i": "1",
		"o": "0",
		"s": "5",
		"t": "7",
	}

	r := rand.New(rand.NewSource(int64(randSeed)))

	input = strings.ToLower(input)

	var result strings.Builder
	for _, ch := range input {
		if r.Float64() < replPerc {
			if replacement, ok := replacements[string(ch)]; ok {
				result.WriteString(replacement)
				continue
			}
		}
		chStr := string(ch)
		if r.Intn(2) == 0 {
			chStr = strings.ToUpper(chStr)
		}
		result.WriteString(chStr)
	}

	return result.String()
}

func main() {
	replacementPerc := flag.Float64("p", 1.0, "Percentage of characters to replace with l33tspeak")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + " "
	}

	fmt.Println(leetSpeak(input, *replacementPerc, time.Now().UnixNano()))
}
