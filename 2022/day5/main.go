package main

import (
	"fmt"
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

// Hardcoded parsing follows after inspecting the input file.
// First 8 lines are the stacks
const stackCount = 9

func main() {
	s := input.OpenFileBuffered("input")

	c := 0

	var stacks []stack = make([]stack, stackCount)
	for _, stack := range stacks {
		stack.crates = make([]rune, 0, 20)
	}

	for s.Scan() {
		line := s.Text()
		fmt.Println(line)

		if c < 8 {
			// Create stack data structure
			crates := readCrateLine([]rune(line))
			for i, r := range crates {
				if r == ' ' {
					fmt.Println("Empty crate")
				} else {
					stacks[i].crates = append(stacks[i].crates, r)
				}
			}
		}

		if c > 8 {
			// Read commands, alter stacks.
			fmt.Println(stacks)
			break
		}

		c++
	}

	// Part1: Print top crate of each stack
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
