package main

import (
	"fmt"

	"github.com/Flonka/advent-of-code-2021/day2"
	"github.com/Flonka/advent-of-code-2021/input"
)

func main() {

	lines := input.ReadLinesInFile("input")

	sum := [2]int{0, 0}

	for _, commandString := range lines {
		c := day2.NewCommandFromInputLine(commandString)

		sum[0] += c.Direction[0]
		sum[1] += c.Direction[1]
	}

	fmt.Println(sum[0] * sum[1])

}
