package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"github.com/Flonka/advent-of-code-2021/input"
)

func main() {

	r := input.OpenFile("ex_input")
	defer r.Close()
	s := bufio.NewScanner(r)

	easyDigitCount := 0
	for s.Scan() {
		l := s.Text()
		d := strings.Split(l, "|")
		signalPatterns := strings.Fields(d[0])
		sort.Slice(signalPatterns, func(i, j int) bool {
			return len(signalPatterns[i]) < len(signalPatterns[j])
		})

		fmt.Println(signalPatterns)

		// for i := 0; i < 10; i++ {

		// }

		outputDigits := strings.Fields(d[1])
		for i := 0; i < 4; i++ {
			switch len(outputDigits[i]) {
			case 2, 3, 4, 7:
				easyDigitCount++
			}
		}
	}

	fmt.Println("Part1:", easyDigitCount)
}

type InputPattern struct {
	input string
}
