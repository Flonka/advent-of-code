package main

import (
	"fmt"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	lines := input.ReadLinesInFile("input.txt")

	timeStrings := strings.Fields(lines[0])[1:]
	times := input.StringsToInts(timeStrings)
	distanceStrings := strings.Fields(lines[1])[1:]
	distances := input.StringsToInts(distanceStrings)

	r1 := 1

	for i := 0; i < len(times); i++ {
		time := times[i]
		recordDistance := distances[i]

		r1 *= bruteForcePart1(time, recordDistance)
	}

	fmt.Println("p1", r1)
}

// Return number of ways to beat record distance
func bruteForcePart1(time int, recordDistance int) int {

	beatCount := 0

	for pressedTime := 1; pressedTime < time; pressedTime++ {

		d := pressedTime * (time - pressedTime)

		if d > recordDistance {
			beatCount++
		}

	}

	// fmt.Println(time, recordDistance, beatCount)

	return beatCount
}
