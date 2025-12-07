package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	lines := input.ReadLinesInFile("input.txt")

	dial := 50
	zcount := 0
	for _, l := range lines {
		t := parseTurn(l)

		dial = dial + t
		if dial >= 100 {
			dial = dial - 100
		} else if dial < 0 {
			dial = dial + 100
		}

		if dial == 0 {
			zcount++
		}
	}

	fmt.Println("Part1", zcount)
}

func parseTurn(s string) int {
	mult := 1
	if s[0] == 'L' {
		mult = -1
	}

	i, err := strconv.Atoi(s[1:])
	if err != nil {
		fmt.Println("failed to parse", s)
		os.Exit(1)
	}

	for i > 100 {
		i -= 100
	}

	return mult * i
}
