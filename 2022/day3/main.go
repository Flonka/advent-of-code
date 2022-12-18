package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	s := input.OpenFileBuffered("input")

	var sum int
	for s.Scan() {
		line := s.Text()
		l := len(line)
		c1 := line[:(l / 2)]
		c2 := line[(l / 2):]

		ps := getCommonPriorities(c1, c2)

		fmt.Println(ps)
		for k := range ps {
			sum += k	
		}
	}


	fmt.Println("task1", sum)

}

func getCommonPriorities(c1 string, c2 string) map[int]int {

	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	var prios map[int]int = make(map[int]int)
	for _, v := range c1 {

		if strings.ContainsRune(c2, v) {

			prio := runeToPrio(v)
			prios[prio]++
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
