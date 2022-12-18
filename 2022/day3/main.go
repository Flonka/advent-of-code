package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	s := input.OpenFileBuffered("input")

	var prioMap map[int]bool = make(map[int]bool)
	for s.Scan() {
		line := s.Text()
		l := len(line)
		c1 := line[:(l / 2)]
		c2 := line[(l / 2):]

		ps := getCommonPriorities(c1, c2)

		for _, i := range ps {
			prioMap[i] = true
		}
	}

	fmt.Println(prioMap)

	var sum int
	for k := range prioMap {
		sum += k
	}

	fmt.Println("task1", sum)

}

func getCommonPriorities(c1 string, c2 string) []int {

	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	var prios []int
	for _, v := range c1 {

		if strings.ContainsRune(c2, v) {

			prio := runeToPrio(v)
			prios = append(prios, prio)
		}

	}

	return prios

}

func runeToPrio(r rune) int {

	var priodiff int
	if unicode.IsUpper(r) {
		// 65 -27 =
		priodiff = 38
	} else {
		// 97 - 1
		priodiff = 96
	}

	prio := int(r) - priodiff
	return prio
}
