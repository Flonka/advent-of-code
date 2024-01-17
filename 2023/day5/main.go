package main

import (
	"fmt"
	"slices"
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

	locs := []int{}
	for _, seed := range seeds {
		locs = append(locs, FindLocation(seed, rangeMaps))
	}

	slices.Sort(locs)
	fmt.Println("Part1", locs[0])

	// part2(seeds, rangeMaps)

}

func part2(seeds []int, maps []RangeMap) {

	minLoc := -1

	for i := 0; i < len(seeds)-1; i += 2 {
		start := seeds[i]
		end := start + seeds[i+1]
		for j := start; j <= end; j++ {
			loc := FindLocation(j, maps)
			if loc < minLoc {
				minLoc = loc
			}
		}
	}

	fmt.Println("Part2", minLoc)
}

func FindLocation(seed int, maps []RangeMap) int {

	fmt.Println("seed", seed)
	// Go through maps in order to find the location
	for i := 0; i < len(maps); i++ {
		m := maps[i]
		seed = SeedTranslation(seed, m)
		// fmt.Println("map", i+1, "seed", seed)
	}

	return seed
}

func SeedTranslation(seed int, m RangeMap) int {
	for j := 0; j < len(m.Entries); j++ {
		e := m.Entries[j]

		// Find if seed is contained between the entry's source and destination range
		if seed >= e.Source && seed <= e.Source+e.Range {
			// fmt.Println("seed found", j)
			diff := e.Destination - e.Source
			return seed + diff
		}
	}
	// Defaults to original value
	return seed
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
