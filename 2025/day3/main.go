package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	lines := input.ReadLinesInFile("input.txt")
	joltage := 0
	for _, l := range lines {
		bank := input.StringsToInts(strings.Split(l, ""))
		max := slices.Max(bank)
		mIdx := slices.Index(bank, max)

		var max2 int
		if mIdx == len(bank)-1 {
			// Find the second largest and update the max to that index and value
			max = slices.Max(bank[:mIdx])
			mIdx = slices.Index(bank, max)
			max = bank[mIdx]
		}
		// Find the second max value
		tail := bank[mIdx+1:]
		max2 = slices.Max(tail)

		js := strconv.Itoa(max) + strconv.Itoa(max2)

		jolt, err := strconv.Atoi(js)
		if err != nil {
			fmt.Println(js, "failed")
		}
		joltage += jolt

	}

	fmt.Println("Joltage", joltage)
}
