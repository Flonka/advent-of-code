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

	bonusCards := make([]int, 0, 100)
	cardCount := 0

	cc := 0
	for s.Scan() {
		fmt.Println("loop", cc)

		line := s.Text()
		c := CardFromLine(line)
		matches := c.GetMatchCount()

		// One for the looped card, plus one for each of the bonuses active
		newCards := 1 + len(bonusCards)*matches
		part1 += part1Score(matches)

		// Add new bonuscards from matches
		bonusCards = updateBonusCards(bonusCards)
		fmt.Println("matches", matches)
		fmt.Println("length before", len(bonusCards))
		for i := 0; i < newCards; i++ {
			bonusCards = append(bonusCards, matches)
		}
		fmt.Println("length after", len(bonusCards))
		cardCount += newCards
		if cc > 5 {
			break
		}
		cc++
		fmt.Println()
	}

	fmt.Println("part1", part1)
	fmt.Println("part2", cardCount)
}

func updateBonusCards(bonusCards []int) []int {
	// Reduce current active bonuses for current line
	for i := 0; i < len(bonusCards); i++ {
		bonusCards[i] = bonusCards[i] - 1
	}
	slices.Sort(bonusCards)
	for i := 0; i < len(bonusCards); i++ {
		// Find first positive number
		n := bonusCards[i]
		if n > 0 {
			// fmt.Println("before", bonusCards)
			bonusCards = slices.Delete(bonusCards, 0, i)
			// fmt.Println("after", bonusCards)
			return bonusCards
		}
	}

	return bonusCards
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
