package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	s := input.OpenFileBuffered("input.txt")
	p1Sum := 0

	for s.Scan() {

		data := input.StringsToInts(strings.Fields(s.Text()))
		p1Sum += part1(data)
	}

	fmt.Println("Part1", p1Sum)
}

// Return extrapolated next value of input data.
func part1(data []int) int {

	diffs := make([][]int, 0)
	diffs = append(diffs, data)
	c := 0
	for {

		ds := createDiffSlice(diffs[c])
		diffs = append(diffs, ds)
		// Compact modifies input slice..
		if len(slices.Compact(slices.Clone(ds))) == 1 {
			break
		}
		c++
	}

	ext := 0
	for i := 0; i < len(diffs); i++ {
		ds := diffs[i]
		// Add last value of arrays togehter
		// if len(ds) != 1 {
		ext += ds[len(ds)-1]
		// }
	}

	return ext
}

func createDiffSlice(in []int) []int {
	out := make([]int, 0, len(in)-1)
	for i := 0; i < len(in)-1; i++ {
		n0 := in[i]
		n1 := in[i+1]

		out = append(out, n1-n0)
	}

	return out
}
