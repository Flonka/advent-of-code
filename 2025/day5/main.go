package main

import (
	"log"
	"log/slog"
	"strconv"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/input"
)

func main() {
	cli.Default()
	ranges, ids := readData()
	slog.Info("Read data lengths", "ranges", len(ranges), "ids", len(ids))

	freshCount := 0

	for _, id := range ids {
		if isFresh(ranges, id) {
			freshCount++
		}
	}

	slog.Info("Problem 1", "freshCount", freshCount)
}

func isFresh(ranges []input.Range, id int) bool {
	for _, r := range ranges {
		if id >= r.Start && id <= r.End {
			return true
		}
	}
	return false
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
