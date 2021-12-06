package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Flonka/advent-of-code-2021/input"
)

func main() {

	lines := input.ReadLinesInFile("input")

	bitCount := [12]int{}

	for _, binaryString := range lines {
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
	var epsilonRateBinString string
	for i := 0; i < len(bitCount); i++ {
		if bitCount[i] > halfAmount {
			gammaRateBinString += "1"
			epsilonRateBinString += "0"
		} else {
			gammaRateBinString += "0"
			epsilonRateBinString += "1"
		}
	}

	fmt.Println(gammaRateBinString)
	fmt.Println(epsilonRateBinString)
	gammaRate, err := strconv.ParseInt(gammaRateBinString, 2, 13)

	if err != nil {
		log.Fatal(err)
	}
	// epsilonRate, err := strconv.ParseInt(epsilonRateBinString, 2, 13)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ones := int(0b_111111111111)
	// epsilon rate should be bitwise inverse of gammrate
	epsilonRate := ones ^ int(gammaRate)

	fmt.Printf("%v\n", gammaRate)
	fmt.Printf("%012b\n", gammaRate)

	fmt.Printf("%v\n", epsilonRate)
	fmt.Printf("%012b\n", epsilonRate)

	fmt.Println("Power consumption:", gammaRate*int64(epsilonRate))

}
