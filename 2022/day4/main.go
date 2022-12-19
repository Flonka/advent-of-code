package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

type assignmentMask = []uint

type assignmentRange struct {
	start uint
	end   uint
}

func main() {
	s := input.OpenFileBuffered("input")

	var fullyCointained int
	for s.Scan() {
		line := s.Text()
		// Get pairs
		pairs := strings.Split(line, ",")

		// Find out how many assignment masks are needed
		// based on max range and intsize of architecture
		prange := make([]assignmentRange, len(pairs))
		for i, v := range pairs {
			s, e := getBounds(v)
			prange[i].start = s
			prange[i].end = e
		}

		// Create bitmasks representing pairs
		masks := createMasks(prange)

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

func printMasks(masks []assignmentMask) {
	// TODO: Print Intsize
	for _, assignment := range masks {
		for i := len(assignment) - 1; i >= 0; i-- {
			fmt.Printf("%064b", assignment[i])
		}
		fmt.Println()
	}

	// fmt.Printf("AND \t %010b\n", masks[0]&masks[1])
	// fmt.Printf("OR \t %010b\n", masks[0]|masks[1])
	// fmt.Printf("XOR \t%010b\n", masks[0]^masks[1])
	// fmt.Printf("NOT \t%010b\n", ^masks[0])
	// fmt.Printf("NOT \t%010b\n", ^masks[1])
	// fmt.Printf("ANDNOT \t%010b\n", masks[0]&^masks[1])
}

// createMasks creates assignmentMask slice from given assignmentRange slice
func createMasks(pairs []assignmentRange) []assignmentMask {

	var maxBound uint
	for _, v := range pairs {
		if v.end > maxBound {
			maxBound = v.end
		}
	}

	// Calcualte needed mask sub slices
	fmt.Println("maxbound", maxBound)
	d := float64(maxBound) / strconv.IntSize
	fmt.Println("d", d)
	subSliceCount := uint(math.Ceil(d))
	fmt.Println("subslicecount", subSliceCount)

	// equal number of masks needed as pairs
	masks := make([]assignmentMask, len(pairs))

	for maskI, v := range pairs {

		masks[maskI] = make(assignmentMask, subSliceCount)

		// assignmentRange is not 1-indexed, so reduce by one for the
		// 0index bit mask, in assignmentmasks
		for i := v.start - 1; i <= v.end-1; i++ {
			// Find correct subslice to set bit on
			// based on i
			subIndex := int(math.Floor(float64(i) / strconv.IntSize))
			// bit position to set  need to be relative to subslice position
			pos := i - (uint(subIndex) * strconv.IntSize)
			// if subIndex > 0 {
			// 	fmt.Println(subIndex, pos)
			// }

			vv := masks[maskI][subIndex]
			masks[maskI][subIndex] = setBit(vv, pos)
			// fmt.Printf("%064b\n", masks[maskI][subIndex])
		}

	}

	return masks

}

// Return start and end of a pair span e.g. 3-7
func getBounds(s string) (uint, uint) {
	spans := strings.Split(s, "-")
	start, err := strconv.Atoi(spans[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(spans[1])
	if err != nil {
		log.Fatal(err)
	}
	return uint(start), uint(end)
}

// Sets the bit at pos in the integer n.
func setBit(n uint, pos uint) uint {
	n |= (1 << pos)
	return n
}

// Returns whether any assignment is contained in
// any of the other assignments
func isContained(assignments []assignmentMask) bool {

	// Loop assignments over each other, comparing all with all
	for i, assignmentA := range assignments {
		for j, assignmentB := range assignments {
			// Skip comparing same assignmentMask
			if j == i {
				continue
			}

			// Compare assignments, each mask has potentially N uint masks
			// lengths of assignmentmasks must be the same.
			neededForContained := len(assignmentA)
			var count int
			for i := 0; i < neededForContained; i++ {
				maskA := assignmentA[i]
				maskB := assignmentB[i]
				if (maskA & maskB) == maskA {
					count++
				}
			}

			if count == neededForContained {
				return true
			}
		}
	}

	return false
}
