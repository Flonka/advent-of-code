package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/input"
)

func main() {

	cli.Default()

	s := input.OpenFileBuffered("input.txt")
	sum := 0
	sum2 := 0
	for s.Scan() {
		l := s.Text()
		n1 := getNumberFromLine(l, false)
		n2 := getNumberFromLine(l, true)
		sum += n1
		sum2 += n2
		fmt.Println()
	}

	slog.Info("Part 1", "sum", sum)
	slog.Info("Part 2", "sum", sum2)

}

var textNumbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getNumberFromLine(line string, includeTextNumbers bool) int {

	line = strings.ToLower(line)
	// textInts , lookup table for each index in the string
	textInts := make([]int, utf8.RuneCountInString(line))
	if includeTextNumbers == true {
		for numberValue, textNumber := range textNumbers {
			// Find all occurences of number
			for _, i := range findAllIndices(line, textNumber) {
				textInts[i] = numberValue
			}
		}
		slog.Debug("Text numbers", "textInts", textInts, "line", line)
	}

	ints := make([]string, 0)
	for i, v := range line {
		if unicode.IsDigit(v) {
			ints = append(ints, string(v))
		}

		if includeTextNumbers == true {
			tInt := textInts[i]
			if tInt > 0 {
				ints = append(ints, fmt.Sprint(tInt))
			}
		}
	}

	slog.Debug("Found integers", "ints", ints)

	combined := fmt.Sprintf("%v%v", ints[0], ints[len(ints)-1])
	i, err := strconv.Atoi(combined)
	slog.Debug("Produced number", "combined", combined, "line", line)
	if err != nil {
		slog.Error("Not an integer", "combined", combined)
		os.Exit(1)
	}

	return i
}

// findAllIndices, return slice of all indices where substring in s
func findAllIndices(s string, substring string) []int {

	indices := make([]int, 0)
	result := strings.Index(s, substring)
	offset := 0
	for result >= 0 {
		indices = append(indices, result+offset)
		// Step one character forwards
		step := result + 1
		s = s[step:]
		// offset needs to be incremented with the same step
		offset += step
		result = strings.Index(s, substring)
	}
	return indices
}
