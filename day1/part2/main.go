package main

import (
	"fmt"

	"github.com/Flonka/advent-of-code-2021/day1"
)

func main() {
	depthData := day1.GetDepthData("input")

	lastD := depthData[0:3]
	lastSum := sum3(lastD)
	incs := 0
	for i := range depthData[1 : len(depthData)-2] {

		sumD := sum3(depthData[i+1 : i+1+3])
		if sumD > lastSum {
			incs++
		}
		lastSum = sumD
	}

	fmt.Println(incs)
}

func sum3(s []int) int {
	return s[0] + s[1] + s[2]
}
