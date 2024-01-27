package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	s := input.OpenFileBuffered("input.txt")

	s.Scan()
	instructions := toInts(s.Text())

	m := readNodeMap(s)

	fmt.Println("Part1", part1(m, instructions))
	fmt.Println("Part2", part2(m, instructions))
}

func part2(m map[string][2]string, instructions []int) int {

	// Find all starting ids
	ids := make([]string, 0)

	for k := range m {
		if k[2] == 'A' {
			ids = append(ids, k)
		}
	}

	count := 0

mainLoop:
	for {
		for i := 0; i < len(ids); i++ {
			ids[i] = runInstructions(ids[i], m, instructions)
		}
		count++

		for i := 0; i < len(ids); i++ {
			if ids[i][2] != 'Z' {
				continue mainLoop
			}
		}
		break
	}

	fmt.Println(ids, count)
	return count * len(instructions)
}

func runInstructions(id string, m map[string][2]string, instructions []int) string {
	for i := 0; i < len(instructions); i++ {
		id = m[id][instructions[i]]
	}

	return id
}

func part1(m map[string][2]string, instructions []int) int {
	// Run instructions until return ZZZ
	count := 0
	// Start at AAA
	id := "AAA"
	for {
		id = runInstructions(id, m, instructions)
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
