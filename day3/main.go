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
	// fmt.Println(bits)
	// fmt.Printf("%0"+fmt.Sprint(bits)+"b\n", criteria)
	// fmt.Printf("%012b\n", 1<<0)
	// fmt.Printf("%012b\n", criteria&(1<<0))
	// fmt.Printf("%v\n", criteria&(1<<0))

	var out []int
	out = data
	bits := 12
	for i := bits - 1; i >= 0; i-- {
		if len(out) < 10 {
			fmt.Println("---")
			for _, v := range out {
				fmt.Printf("%012b\n", v)
			}
			fmt.Println("---")
		}
		out = filterInts(out, i, most)
		fmt.Println("After filter", len(out))

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
	var setBitCount int
	mask := 1 << pos
	for _, v := range data {
		x := v & mask
		// fmt.Printf("%012b value\n", v)
		// fmt.Printf("%012b mask\n", mask)
		// fmt.Printf("%012b x %v\n", x, x)
		x = x >> pos
		setBitCount += x
	}
	// Create criteria int , 1 or 0 depending on most common.
	var criteria int
	half := len(data) / 2

	if setBitCount == len(data) {
		criteria = 1
	} else if setBitCount >= half && mostCommon {
		criteria = 1
	}

	return criteria << pos
}

func filterInts(data []int, pos int, mostCommon bool) []int {
	// fmt.Println("filter", pos)
	mask := 1 << pos
	criteria := bitmaskCriteria(data, pos, mostCommon)
	var filtered []int = make([]int, 0)
	for _, v := range data {
		maskedV := v & mask
		if maskedV == criteria {
			filtered = append(filtered, v)
			if len(data) < 10 {
				fmt.Printf("%012b v \n", v)
				fmt.Printf("%012b mask\n", mask)
				fmt.Printf("%012b masked V \n", maskedV)
				fmt.Printf("%012b criteria \n", criteria)
				fmt.Printf("%012b added\n", v)
			}
		} else {
			if len(data) < 10 {
				fmt.Printf("%012b v \n", v)
				fmt.Printf("%012b mask\n", mask)
				fmt.Printf("%012b masked V \n", maskedV)
				fmt.Printf("%012b criteria \n", criteria)
				fmt.Printf("%012b removed\n", v)
			}
			// fmt.Printf("%012b masked C \n", maskedCrit)
			// fmt.Printf("%012b result %v\n", comp, maskedCrit == maskedV)
		}

	}

	return filtered
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
