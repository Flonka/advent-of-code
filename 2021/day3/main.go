package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Flonka/advent-of-code-2021/input"
)

func main() {

	allInts := powerConsumption()

	oxygenGeneratorRating := findByCriteria(allInts, true)
	c02ScrubberRating := findByCriteria(allInts, false)
	fmt.Println("Life Support Rating:", oxygenGeneratorRating*c02ScrubberRating)
}

func powerConsumption() []int {
	lines := input.ReadLinesInFile("input")

	allInts := make([]int, 0, len(lines))

	bitCount := [12]int{}

	for _, binaryString := range lines {

		lineInt, err := strconv.ParseInt(binaryString, 2, 16)
		if err != nil {
			log.Fatal(err)
		}
		allInts = append(allInts, int(lineInt))

		for i := 0; i < len(bitCount); i++ {

			bitvalue, err := strconv.Atoi(string(binaryString[i]))
			if err != nil {
				log.Fatal(err)
			}

			bitCount[i] += bitvalue
		}
	}

	halfAmount := len(lines) / 2
	var gammaRateBinString string
	for i := 0; i < len(bitCount); i++ {
		if bitCount[i] >= halfAmount {
			gammaRateBinString += "1"
		} else {
			gammaRateBinString += "0"
		}
	}

	gammaRate, err := strconv.ParseInt(gammaRateBinString, 2, 0)

	if err != nil {
		log.Fatal(err)
	}

	ones := int(0b_111111111111)
	// epsilon rate should be bitwise inverse of gammrate
	epsilonRate := ones ^ int(gammaRate)

	// fmt.Printf("%v\n", gammaRate)
	// fmt.Printf("%012b\n", gammaRate)

	// fmt.Printf("%v\n", epsilonRate)
	// fmt.Printf("%012b\n", epsilonRate)

	fmt.Println("Power Consumption:", gammaRate*int64(epsilonRate))

	return allInts

}

func findByCriteria(data []int, most bool) int {
	var out []int
	out = data
	bits := 12
	for i := bits - 1; i >= 0; i-- {

		out = filterInts(out, i, most)

		if len(out) == 1 {
			return out[0]
		}
	}

	log.Fatal(out)
	return out[0]
}

// Return bitcriteria to use in filtering out values
func bitmaskCriteria(data []int, pos int, mostCommon bool) int {

	// Find count of set bit by position
	var setBits int
	mask := 1 << pos
	for _, v := range data {
		x := v & mask
		x = x >> pos
		setBits += x
	}
	// Create criteria int , 1 or 0 depending on most common.
	unsetBits := len(data) - setBits

	if setBits == len(data) {
		return 1 << pos
	}

	if setBits == 0 {
		return 0
	}

	var criteria int
	if mostCommon {
		if setBits >= unsetBits {
			criteria = 1
		}
	} else if setBits < unsetBits {
		criteria = 1
	}

	return criteria << pos
}

func filterInts(data []int, pos int, mostCommon bool) []int {
	mask := 1 << pos
	criteria := bitmaskCriteria(data, pos, mostCommon)
	var filtered []int = make([]int, 0)
	for _, v := range data {
		maskedV := v & mask
		if maskedV == criteria {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
