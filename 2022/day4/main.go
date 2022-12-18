package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	s := input.OpenFileBuffered("input_test")

	var fullyCointained int
	for s.Scan() {
		line := s.Text()
		// Get pairs
		pairs := strings.Split(line, ",")

		// Create bitmasks representing pairs
		masks := createMasks(pairs)

		// Check if they are fully cointained with eachother
		// and increment counter
		for _, m := range masks {
			fmt.Printf("%010b\n", m)
		}
	}

	fmt.Println("Part1:", fullyCointained)
}

func createMasks(pairs []string) []int {

	fmt.Println(pairs)

	masks := make([]int, len(pairs))

	for maskI, v := range pairs {
		spans := strings.Split(v, "-")
		start, _ := strconv.Atoi(spans[0])
		end, _ := strconv.Atoi(spans[1])
		fmt.Println(start, end)

		for i := start; i <= end; i++ {
			masks[maskI] = setBit(masks[maskI], i-1)
		}

	}

	return masks

}

// Sets the bit at pos in the integer n.
func setBit(n int, pos int) int {
	n |= (1 << pos)
	return n
}
