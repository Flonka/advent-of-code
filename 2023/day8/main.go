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

	// Convert inputLine to ints
	m := readNodeMap(s)

	// Run instructions until return ZZZ
	count := 0
	// Start at AAA
	id := "AAA"
	for {
		for i := 0; i < len(instructions); i++ {
			id = m[id][instructions[i]]
		}
		count++

		if id == "ZZZ" {
			break
		}
	}

	fmt.Println("Part1", count*len(instructions))

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
