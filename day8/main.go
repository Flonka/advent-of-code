package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/Flonka/advent-of-code-2021/input"
)

func main() {

	readData()
}

func readData() {
	r := input.OpenFile("input")
	defer r.Close()
	s := bufio.NewScanner(r)

	easyDigitCount := 0
	for s.Scan() {
		l := s.Text()
		d := strings.Split(l, "|")
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
