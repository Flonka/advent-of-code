package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Flonka/advent-of-code-2021/input"
)

func main() {

	lines := input.ReadLinesInFile("input")

	allInts := make([]int, 0, len(lines))

	bitCount := [12]int{}

	for _, binaryString := range lines {

		lineInt, err := strconv.Atoi(binaryString)
		if err != nil {
			log.Fatal(err)
		}
		allInts = append(allInts, lineInt)

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
		if bitCount[i] > halfAmount {
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

	fmt.Printf("%v\n", gammaRate)
	fmt.Printf("%012b\n", gammaRate)

	fmt.Printf("%v\n", epsilonRate)
	fmt.Printf("%012b\n", epsilonRate)

	fmt.Println("Power Consumption:", gammaRate*int64(epsilonRate))

	findByCriteria(allInts, int(gammaRate), len(bitCount))
	// oxygenGeneratorRating := findByCriteria(allInts, int(gammaRate))
	// c02ScrubberRating := findByCriteria(allInts, epsilonRate)
	// fmt.Println("Life Support Rating:", oxygenGeneratorRating*c02ScrubberRating)
}

func findByCriteria(data []int, criteria int, bits int) int {
	fmt.Println(bits)
	fmt.Printf("%0"+fmt.Sprint(bits)+"b\n", criteria)
	fmt.Printf("%012b\n", 1<<0)
	fmt.Printf("%012b\n", criteria&(1<<0))
	fmt.Printf("%v\n", criteria&(1<<0))

	out := make([]int, 0, len(data))
	out = data
	for i := bits - 1; i >= 0; i-- {
		out = filterInts(out, criteria, i)
		if len(out) == 1 {
			return out[0]
		}
	}

	// log.Fatal(out)
	return 0
}

func filterInts(data []int, criteria int, pos int) []int {
	for _, v := range data {
		fmt.Println(v)
	}

	return data
}

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
