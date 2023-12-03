package main

import (
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
		sum += n
	}

	if sum != 281 {
		t.Fatal("Sum not correct")
	}
}

func TestMultipleTextNumbers(t *testing.T) {

	test := "4onefive6zsjhzvrjnsfive1five"

	n := getNumberFromLine(test, true)
	if n != 45 {
		t.Errorf("Number %v != %v", n, 45)
	}

}
