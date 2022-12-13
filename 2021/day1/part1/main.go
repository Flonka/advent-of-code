package main

import (
	"fmt"

	"github.com/Flonka/advent-of-code-2021/day1"
)

func main() {

	depthData := day1.GetDepthData("input")

	lastD := depthData[0]
	incs := 0
	for _, d := range depthData[1:] {

		if d > lastD {
			incs++
		}
		lastD = d
	}

	fmt.Println(incs)
}
