package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	s := input.OpenFileBuffered("input")

	var fullyCointained int
	for s.Scan() {
		line := s.Text()
		// Get pairs
		pairs := strings.Split(line, ",")

		// Create bitmasks representing pairs
		masks := createMasks(pairs)

		// Check if they are fully cointained with eachother
		// and increment counter
		if isContained(masks) {
			fmt.Println(pairs)
			printMasks(masks)
			fmt.Println("Contained")
			fullyCointained++
			fmt.Println()
		}

	}

	fmt.Println("Part1:", fullyCointained)
}

func printMasks(masks []uint) {
	fmt.Printf("%0100b\n%0100b\n", masks[0], masks[1])

	// fmt.Printf("AND \t %010b\n", masks[0]&masks[1])
	// fmt.Printf("OR \t %010b\n", masks[0]|masks[1])
	// fmt.Printf("XOR \t%010b\n", masks[0]^masks[1])
	// fmt.Printf("NOT \t%010b\n", ^masks[0])
	// fmt.Printf("NOT \t%010b\n", ^masks[1])
	// fmt.Printf("ANDNOT \t%010b\n", masks[0]&^masks[1])
}

func createMasks(pairs []string) []uint {

	masks := make([]uint, len(pairs))

	for maskI, v := range pairs {
		spans := strings.Split(v, "-")
		start, _ := strconv.Atoi(spans[0])
		end, _ := strconv.Atoi(spans[1])

		for i := start; i <= end; i++ {
			masks[maskI] = setBit(masks[maskI], uint(i-1))
		}

	}

	return masks

}

// Sets the bit at pos in the integer n.
func setBit(n uint, pos uint) uint {
	n |= (1 << pos)
	return n
}

// Returns whether any assignment is contained in
// any of the other assignments
func isContained(assignments []uint) bool {
	for i, a := range assignments {

		for j, b := range assignments {
			if j == i {
				continue
			}

			if (a & b) == a {
				return true
			}
		}
	}

	return false
}
