package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	d := readInts()

	fmt.Println(d[0])

}

func readInts() [][]int {

	rows := make([][]int, 0, 100)
	s := input.OpenFileBuffered("input")

	for s.Scan() {
		ints := make([]int, 0, 100)

		line := s.Text()

		for _, v := range strings.Split(line, "") {
			i, _ := strconv.Atoi(v)
			ints = append(ints, i)
		}
		rows = append(rows, ints)
	}

	return rows
}
