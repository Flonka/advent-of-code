package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/mathutils"
)

func main() {

	s := input.OpenFileBuffered("input.txt")

	s.Scan()
	instructions := toInts(s.Text())

	m := readNodeMap(s)

	fmt.Println("Part1", part1(m, instructions))
	fmt.Println("Part2", part2(m, instructions))
}

func part2(m map[string][2]string, ins []int) int {

	startIds := part2Starts(m, ins)
	if !verifyCycling(startIds, m, ins) {
		fmt.Println("Aint no cycling")
		os.Exit(1)
	}

	// Cycling with consequent loop iterations, the LCM method will work.

	// Get all iterations
	cs := make([]int, 0, len(startIds))
	for i := 0; i < len(startIds); i++ {
		_, c := findGoal(startIds[i], m, ins)
		cs = append(cs, c)
	}

	lcm := mathutils.LCM(cs[0], cs[1], cs[2:]...)

	return lcm * len(ins)
}

func verifyCycling(starts []string, m map[string][2]string, instructions []int) bool {

	// need to be producing the same iterations between finding the goal
	// as well as identifying looping (start == end)
	// Look for cycles
	for i := 0; i < len(starts); i++ {

		id := starts[i]
		counts := make([]int, 0)
		for {
			startId := id
			var c int
			fmt.Println("start", id)
			id, c = findGoal(id, m, instructions)
			counts = append(counts, c)

			if startId == id {
				// Start == end, its looping.
				// All counts must be the same..
				counts = slices.Compact(counts)
				if len(counts) != 1 {
					return false
				}
				break
			}
		}
	}

	return true

}

func part2Starts(m map[string][2]string, instructions []int) []string {
	// Find all starting ids
	ids := make([]string, 0)

	for k := range m {
		if k[2] == 'A' {
			ids = append(ids, k)
		}
	}

	return ids
}

func findGoal(id string, m map[string][2]string, instructions []int) (string, int) {
	count := 0
	for {
		id = runInstructions(id, 1, m, instructions)
		count++

		if id[2] == 'Z' {
			break
		}

	}
	return id, count
}

func runInstructions(id string, times int, m map[string][2]string, instructions []int) string {
	for x := 0; x < times; x++ {
		for i := 0; i < len(instructions); i++ {
			id = m[id][instructions[i]]
		}

	}

	return id
}

func part1(m map[string][2]string, instructions []int) int {
	// Run instructions until return ZZZ
	count := 0
	// Start at AAA
	id := "AAA"
	for {
		id = runInstructions(id, 1, m, instructions)
		count++

		if id == "ZZZ" {
			break
		}
	}

	return count * len(instructions)
}

func toInts(s string) []int {

	ints := make([]int, 0, len(s))
	for _, v := range s {
		switch v {

		case 'L':
			ints = append(ints, 0)

		case 'R':
			ints = append(ints, 1)
		}
	}

	return ints
}

func readNodeMap(s *bufio.Scanner) map[string][2]string {

	m := make(map[string][2]string, 200)

	for s.Scan() {

		l := s.Text()
		if len(l) == 0 {
			continue
		}

		s1 := strings.Fields(l)
		// fmt.Println(s1[0], s)
		m[s1[0]] = [2]string{s1[2][1:4], s1[3][:3]}
	}

	return m
}
