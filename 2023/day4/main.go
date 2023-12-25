package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

type Card struct {
	WinningNumbers []int
	GottenNumbers  []int
}

func (c *Card) GetMatchCount() int {
	var n int

	for i := 0; i < len(c.GottenNumbers); i++ {
		if slices.Contains(c.WinningNumbers, c.GottenNumbers[i]) {
			n++
		}
	}

	return n
}

func CardFromLine(line string) Card {

	// only use string after colon
	cardString := strings.Split(line, ":")[1]
	s := strings.Split(cardString, "|")

	return Card{
		WinningNumbers: numbersFromString(s[0]),
		GottenNumbers:  numbersFromString(s[1]),
	}
}

func numbersFromString(s string) []int {

	numString := strings.Fields(s)

	nums := make([]int, 0, len(numString))
	for idx, ns := range numString {

		i, err := strconv.Atoi(ns)
		if err != nil {
			fmt.Println(idx, ns, err)
			os.Exit(1)
		}
		nums = append(nums, i)
	}

	return nums
}

func main() {

	s := input.OpenFileBuffered("input.txt")
	part1 := 0

	cardCount := map[int]int{}

	idx := 0
	for s.Scan() {

		line := s.Text()
		c := CardFromLine(line)
		matches := c.GetMatchCount()

		// Part1
		part1 += part1Score(matches)

		// Part2
		// Always one per loop
		cardCount[idx] += 1

		// Add bonuscards
		for i := 1; i <= matches; i++ {
			cardCount[idx+i] += cardCount[idx]
		}

		idx++
	}

	fmt.Println("part1", part1)

	part2 := 0
	for i := 0; i < idx; i++ {
		part2 += cardCount[i]
	}
	fmt.Println("part2", part2)
}

func part1Score(m int) int {

	if m == 0 {
		return 0
	}

	score := 1
	// pow 2 ^ matchount-1
	for i := 0; i < m-1; i++ {
		score = score * 2
	}

	return score
}
