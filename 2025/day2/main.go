package main

import (
	"fmt"

	"github.com/Flonka/advent-of-code/input"
)

func main() {
	rangeS := input.ReadCommaSeparated[string]("input.txt")

	for _, s := range rangeS {
		fmt.Println(s)
	}
}
