package main

import (
	"fmt"

	"github.com/Flonka/advent-of-code-2021/day2"
	"github.com/Flonka/advent-of-code-2021/input"
)

func main() {

	lines := input.ReadLinesInFile("input")

	aim := 0
	depth := 0
	horizontal := 0

	for _, commandString := range lines {
		c := day2.NewCommandFromInputLine(commandString)

		horizontal += c.Direction[0]
		depth += aim * c.Direction[0]
		aim += c.Direction[1]
	}

	fmt.Println(horizontal * depth)

}
