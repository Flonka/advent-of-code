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
	var overlapping int
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
		fmt.Println(pairs)
		printMasks(masks)

		containedFunc := func(a, b uint) bool {
			return (a & b) == a
		}

		overlapFunc := func(a, b uint) bool {
			return a&b > 0
		}

		if bitmaskComp(masks, true, containedFunc) {
			fmt.Println("Contained")
			fullyCointained++
		}
		if bitmaskComp(masks, false, overlapFunc) {
			fmt.Println("Overlapping")
			overlapping++
		}
		fmt.Println()
	}

	fmt.Println("Part1:", fullyCointained)
	fmt.Println("Part2:", overlapping)
}

func printMasks(masks []assignmentMask) {
	// TODO: Print Intsize
	for _, assignment := range masks {
		for i := len(assignment) - 1; i >= 0; i-- {

			fmt.Printf("%064b", assignment[i])
		}
		fmt.Println()
	}
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
	d := float64(maxBound) / strconv.IntSize
	subSliceCount := uint(math.Ceil(d))

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

			vv := masks[maskI][subIndex]
			masks[maskI][subIndex] = setBit(vv, pos)
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

// Returns whether operatorFunc is true for comparing two masks for
// any of the given masks.
// matchAll sets if the operation needs to match all bitmasks between assignments.
func bitmaskComp(assignments []assignmentMask, matchAll bool, operatorFunc func(a, b uint) bool) bool {

	// Loop assignments over each other, comparing all with all
	for i, assignmentA := range assignments {

		var matchCount int = 1
		if matchAll {
			matchCount = len(assignmentA)
		}
		for j, assignmentB := range assignments {
			// Skip comparing same assignmentMask
			if j == i {
				continue
			}

			// Compare assignments, each mask has potentially N uint masks
			// lengths of assignmentmasks must be the same.
			var count int
			for i := 0; i < len(assignmentA); i++ {
				maskA := assignmentA[i]
				maskB := assignmentB[i]
				if operatorFunc(maskA, maskB) {
					count++
					if count == matchCount {
						return true
					}
				}
			}

		}
	}

	return false
}
