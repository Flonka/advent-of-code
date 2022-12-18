package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	s := input.OpenFileBuffered("input")

	var sum int
	var groupSum int
	var lines []string
	var groupPrios map[int]int = make(map[int]int)
	for s.Scan() {
		line := s.Text()
		l := len(line)
		c1 := line[:(l / 2)]
		c2 := line[(l / 2):]
		ps := getCommonPriorities(c1, c2)
		// Task 1
		for k := range ps {
			sum += k
		}

		// Task 2
		lines = append(lines, line)

		if len(lines) == 3 {
			// Every third line, find groupPrio result + reset

			for k := range(getCommonPriorities(lines[0], lines[1])) {
				groupPrios[k]++
			}
			for k := range(getCommonPriorities(lines[0], lines[2])) {
				groupPrios[k]++
			}

			found := false
			for k, v := range groupPrios {
				if v == 2 {
					groupSum += k
					found = true
				} 
			}

			if !found {
				log.Fatal("didnt find it", groupPrios)
			}

			groupPrios = make(map[int]int)
			lines = make([]string, 0, 10)
		}
	}

	fmt.Println("task1", sum)
	fmt.Println("task2", groupSum)

}

func getCommonPriorities(c1 string, c2 string) map[int]int {

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
