package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

type RangeMap struct {
	Entries []RangeEntry
}

type RangeEntry struct {
	Source      int
	Destination int
	Range       int
}

func main() {

	s := input.OpenFileBuffered("input.txt")

	// First line contains the seeds
	s.Scan()
	seeds := GetSeeds(s.Text())

	rangeMaps := make([]RangeMap, 0, 10)
	for s.Scan() {

		line := s.Text()
		if len(line) == 0 {
			continue
		}

		newMap := strings.Contains(line, "map:")
		if newMap {
			fmt.Println("line", line)
			m := RangeMap{
				Entries: make([]RangeEntry, 0),
			}
			rangeMaps = append(rangeMaps, m)
		} else {
			idx := len(rangeMaps) - 1
			m := rangeMaps[idx]
			m.Entries = append(m.Entries, GetEntry(line))
			rangeMaps[idx] = m
		}

	}

	fmt.Println(len(rangeMaps))
	fmt.Println(seeds)
}

func GetEntry(line string) RangeEntry {

	ints := input.StringsToInts(strings.Fields(line))
	r := RangeEntry{
		Destination: ints[0],
		Source:      ints[1],
		Range:       ints[2],
	}

	return r
}

func GetSeeds(line string) []int {

	seeds := []int{}
	seedLine := strings.Replace(line, "seeds:", "", 1)
	for _, strInt := range strings.Fields(seedLine) {
		n, _ := strconv.Atoi(strInt)
		seeds = append(seeds, n)
	}

	return seeds
}
