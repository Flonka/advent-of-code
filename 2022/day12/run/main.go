package main

import (
	"fmt"

	"github.com/Flonka/advent-of-code/2022/day12"
	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

func main() {

	m := readMapData()
	fmt.Println(m)

}

func readMapData() day12.ClimbMap {

	lines := input.ReadLinesInFile("input")

	var w, h int

	// assume all lines are equal
	w = len(lines[0])
	h = len(lines)

	return day12.ClimbMap{
		MapData: spatial.NewDiscreteMap2D(w, h, 2),
	}

}
