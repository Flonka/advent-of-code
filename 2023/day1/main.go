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

	logLevel := cli.Default()

	logLevel.Set(slog.LevelDebug)
	s := input.OpenFileBuffered("input.txt")
	sum := 0
	sum2 := 0
	for s.Scan() {
		l := s.Text()
		sum += getNumberFromLine(l, false)
		sum2 += getNumberFromLine(l, true)
		fmt.Println()
	}

	slog.Info("Part 1", "sum", sum)
	slog.Info("Part 2", "sum", sum2)

}

var textNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getNumberFromLine(line string, includeTextNumbers bool) int {

	line = strings.ToLower(line)
	// textInts , lookup table for each index in the string
	textInts := make([]int, utf8.RuneCountInString(line))
	if includeTextNumbers == true {
		for i, n := range textNumbers {
			iResult := strings.Index(line, n)
			if iResult >= 0 {
				// Found number
				textInts[iResult] = i + 1
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
