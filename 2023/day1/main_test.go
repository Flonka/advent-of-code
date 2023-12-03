package main

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	sum := 0
	for _, v := range input {

		n := getNumberFromLine(v, true)
		fmt.Println(n)
		sum += n
	}

	if sum != 281 {
		t.Fatal("Sum not correct")
	}
}
