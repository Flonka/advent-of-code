package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	lines := input.ReadLinesInFile("input.txt")

	locations := [2][]int{}
	locations[0] = make([]int, 0, 100)
	locations[1] = make([]int, 0, 100)

	for _, line := range lines {

		s := strings.Fields(line)
		i1, err := strconv.Atoi(s[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		i2, err := strconv.Atoi(s[1])

		locations[0] = append(locations[0], i1)
		locations[1] = append(locations[1], i2)
	}

	slices.Sort(locations[0])
	slices.Sort(locations[1])

	if len(locations[0]) != len(locations[1]) {
		fmt.Println("Lenghts do not match")
		os.Exit(1)
	}

	var distance int
	for i := 0; i < len(locations[0]); i++ {

		i1 := locations[0][i]
		i2 := locations[1][i]

		d := i2 - i1

		if d < 0 {
			d = -d
		}

		distance += d
	}

	fmt.Println("Part1, total distance:", distance)

	var similarity int

	for _, n := range locations[0] {
		similarity += n * findN(n, locations[1])
	}

	fmt.Println("Part2, similarity:", similarity)
}

// findN returns count of n in list, assuming list is sorted
func findN(n int, list []int) int {
	var c int
	for _, i := range list {
		if i == n {
			c++
		}
		// early exit due to sorted input list
		if i > n {
			return c
		}
	}

	return c
}
