package main

import (
	"fmt"
	"strings"

	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/mathutils"
)

func main() {

	lines := input.ReadLinesInFile("input.txt")

	var safeCount int

	for _, line := range lines {
		if isReportSafe(input.StringsToInts(strings.Fields(line))) {
			safeCount++
		}
	}

	fmt.Println("Part1, safety count:", safeCount)

}

func isReportSafe(levels []int) bool {
	// A report only counts as safe if both of the following are true:
	//  - The levels are either all increasing or all decreasing.
	//  - Any two adjacent levels differ by at least one and at most three.
	if len(levels) <= 2 {
		return true
	}

	var dt int
	// initialize data from first step
	i1 := levels[0]
	i2 := levels[1]
	dt = i2 - i1
	diff := mathutils.AbsInt(dt)

	if !diffRule(diff) {
		return false
	}

	for i := 2; i < len(levels); i++ {
		i1 := levels[i-1]
		i2 := levels[i]
		dd := i2 - i1

		if dt > 0 && dd < 0 {
			return false
		} else if dt < 0 && dd > 0 {
			return false
		}

		diff := mathutils.AbsInt(dd)
		if !diffRule(diff) {
			return false
		}
	}

	return true
}

func diffRule(diff int) bool {
	if diff >= 1 && diff <= 3 {
		return true
	}
	return false

}
