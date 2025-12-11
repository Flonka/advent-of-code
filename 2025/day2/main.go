package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

type Range struct {
	start int
	end   int
}

// newRangeFromString creates from problem input
// format start-end
func newRangeFromString(input string) Range {
	s := strings.Split(input, "-")

	i1, err := strconv.Atoi(s[0])
	if err != nil {
		os.Exit(1)
	}
	i2, err := strconv.Atoi(s[1])
	if err != nil {
		os.Exit(1)
	}

	return Range{start: i1, end: i2}
}

func main() {
	rangeS := input.ReadCommaSeparated[string]("input.txt")

	sum := 0

	for _, s := range rangeS {
		r := newRangeFromString(s)

		for i := r.start; i <= r.end; i++ {
			if isInvalid(i) {
				sum += i
			}
		}
	}

	fmt.Println("Part1", sum)
}

// isInvalid returns true if the productId is invalid
func isInvalid(productID int) bool {
	if productID > 10 {
		// 1 Convert i to string
		// 2 find repeating pattern in string (must be directly after one another)
		s := strconv.Itoa(productID)

		// Odd length should always be valid due to not repeating two digits.
		if len(s)%2 != 0 {
			return false
		}

		// Go division rounds positive division down ( floor )
		maxL := len(s) / 2

		// string mathcing length span can be max half the length of the string
		matchS := s[0:maxL]
		if strings.Index(s[maxL:], matchS) == 0 {
			fmt.Println("ID:", productID)
			fmt.Println(matchS, "in", s[maxL:])
			return true
		}
	}

	return false
}
