package main

import (
	"fmt"
	"strconv"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	readData()
}

func readData() ([]input.Range, []int) {
	lines := input.ReadLinesInFile("input.txt")
	emptyLine := false

	ids := make([]int, 0, 100)

	for _, l := range lines {
		if len(l) == 0 {
			emptyLine = true
		}
		if emptyLine {
			i, _ := strconv.Atoi(l)
			ids = append(ids, i)
		} else {
			fmt.Println("read range", l)
		}
	}
	return nil, ids
}
