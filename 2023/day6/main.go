package main

import (
	"fmt"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	s := input.OpenFileBuffered("input.txt")

	for s.Scan() {

		line := s.Text()
		fmt.Println(line)
	}
}
