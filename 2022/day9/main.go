package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	s := input.OpenFileBuffered("input")

	for s.Scan() {
		line := s.Text()
		splits := strings.Split(line, " ")

		direction := splits[0]
		amount, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(direction, amount)
	}
}
