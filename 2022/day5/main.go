package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

type stack struct {
	crates []rune
}

func (s stack) String() string {
	b := strings.Builder{}
	for _, v := range s.crates {
		fmt.Fprintf(&b, "%c", v)
	}

	return b.String()
}

type moveCmd struct {
	crateCount int
	fromStack  int
	toStack    int
}

// Hardcoded parsing follows after inspecting the input file.
// First 8 lines are the stacks
const stackCount = 9

func main() {
	s := input.OpenFileBuffered("input")

	c := 0

	var stacks []stack = make([]stack, stackCount)
	for _, stack := range stacks {
		stack.crates = make([]rune, 0, 50)
	}

	for s.Scan() {
		line := s.Text()

		if c < 8 {
			// Create stack data structure
			crates := readCrateLine([]rune(line))
			for i, r := range crates {
				if r != ' ' {
					stacks[i].crates = append(stacks[i].crates, r)
				}
			}
		}

		if c > 9 {
			// Read commands, alter stacks.
			command := readMoveCmd(line)
			// applyCommand(command, stacks)
			applyCommand2(command, stacks)
		}

		c++
	}

	fmt.Print("Topstacks : ")
	for _, v := range stacks {
		fmt.Print(string(v.crates[0]))
	}
	fmt.Println()
}

// Part2 logic
func applyCommand2(c moveCmd, stacks []stack) {
	from := stacks[c.fromStack-1].crates
	to := stacks[c.toStack-1].crates

	to = prependSlice(to, from[:c.crateCount])
	// Remove moved crates
	from = from[c.crateCount:]

	// Update stack crate slices
	stacks[c.toStack-1].crates = to
	stacks[c.fromStack-1].crates = from
}

func prependSlice(x []rune, y []rune) []rune {
	x = append(x, y...)
	copy(x[len(y):], x)
	copy(x[:len(y)], y)
	return x
}

// Part1 logic
func applyCommand(c moveCmd, stacks []stack) {
	from := stacks[c.fromStack-1].crates
	to := stacks[c.toStack-1].crates

	// Move one crate at a time, last pos in crate slice is bottom
	for i := 0; i < c.crateCount; i++ {
		to = prepend(to, from[0])
		from = from[1:]
	}

	// Update stack crate slices
	stacks[c.toStack-1].crates = to
	stacks[c.fromStack-1].crates = from

}

func prepend(x []rune, y rune) []rune {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}

func readMoveCmd(line string) moveCmd {

	s := strings.Split(line, " ")
	c1, err := strconv.Atoi(s[1])
	if err != nil {
		os.Exit(1)
	}
	c2, err := strconv.Atoi(s[3])
	if err != nil {
		os.Exit(1)
	}
	c3, err := strconv.Atoi(s[5])
	if err != nil {
		os.Exit(1)
	}
	return moveCmd{
		crateCount: c1,
		fromStack:  c2,
		toStack:    c3,
	}
}

func readCrateLine(line []rune) []rune {

	runes := []rune{}

	// rune should be from index 2 , every 4 chars
	for i := 1; i < len(line); i += 4 {

		r := line[i]
		runes = append(runes, r)
	}

	return runes
}
