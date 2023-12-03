package main

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/Flonka/advent-of-code/input"
)

func main() {

	s := input.OpenFileBuffered("input.txt")
	for s.Scan() {
		l := s.Text()
		nums := getNumbersFromLine(l)
		fmt.Println(nums)
	}

}

func getNumbersFromLine(line string) []int {

	ints := make([]int, 0)
	for _, v := range line {
		s := string(v)
		i, err := strconv.Atoi(s)
		if err != nil {
			slog.Debug("Not a integer:", "string", s)
		} else {
			ints = append(ints, i)
		}

	}

	return ints
}
