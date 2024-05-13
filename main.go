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

func unleetSpeak(input string) string {
	replacements := map[string]string{
		"4": "a",
		"3": "e",
		"1": "i",
		"0": "o",
		"5": "s",
		"7": "t",
	}

	var result strings.Builder
	for _, ch := range input {
		chStr := string(ch)
		if replacement, ok := replacements[chStr]; ok {
			result.WriteString(replacement)
			continue
		}
		result.WriteString(chStr)
	}

	return result.String()
}

func leetSpeak(input string, ignoreList string, replPerc float64, randSeed int64) string {
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
	var ignoreListSlice []string
	if ignoreList != "" {
		ignoreListSlice = createIgnoreList(ignoreList)
	}
	for _, ch := range input {
		if isOnIgnoreList(string(ch), ignoreListSlice) {
			result.WriteString(string(ch))
			continue
		}
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

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + " "
	}
	return input
}

func createIgnoreList(ignoreList string) []string {
	return strings.Split(ignoreList, ",")
}

func isOnIgnoreList(ch string, ignoreList []string) bool {
	for _, igCh := range ignoreList {
		if igCh == ch {
			return true
		}
	}
	return false
}

func main() {
	replacementPerc := flag.Float64("p", 1.0, "Percentage of characters to replace with l33tspeak")
	isDecoding := flag.Bool("d", false, "Set this flag if you want to decode a l33tspeak string")
	ignoreList := flag.String("ignore", "", "Comma separated list of characters to ignore")
	flag.Parse()

	var input string
	// If we have passed the input string as a command argument, prioritize that. Else use STDIN.
	if len(flag.Args()) > 0 {
		input = flag.Arg(0)
	} else {
		stat, err := os.Stdin.Stat()
		if err != nil {
			fmt.Println("Error checking for STDIN: ", err)
			os.Exit(1)
		}
		// If STDIN is not a pipe or a file, it's probably an interactive input
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			fmt.Println("No input provided. Please provide input via argument or STDIN")
			flag.Usage()
			os.Exit(1)
		}
		input = readInput()
	}

	result := ""
	if *isDecoding {
		result = unleetSpeak(input)
	} else {
		result = leetSpeak(input, *ignoreList, *replacementPerc, time.Now().UnixNano())
	}

	fmt.Println(result)
}
