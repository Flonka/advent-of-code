package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	ranges, ids := readData()
	fmt.Println(len(ranges), len(ids))
}

func readData() ([]input.Range, []int) {
	lines := input.ReadLinesInFile("input.txt")
	emptyLine := false

	ranges := make([]input.Range, 0, 100)
	ids := make([]int, 0, 100)

	for _, l := range lines {
		if len(l) == 0 {
			emptyLine = true
			continue
		}
		if emptyLine {
			i, err := strconv.Atoi(l)
			if err != nil {
				log.Fatalln(err)
			}
			ids = append(ids, i)
		} else {
			r, err := input.NewRangeFromString(l)
			if err != nil {
				log.Fatalln(err)
			}
			ranges = append(ranges, r)
		}
	}
	return ranges, ids
}
