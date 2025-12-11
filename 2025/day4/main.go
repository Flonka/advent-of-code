package main

import (
	"fmt"

	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

func main() {
	dmap := createMap()
	var pos spatial.DiscretePos2D

	rollsLifted := 0

	for w := range dmap.Width {
		pos.X = w
		for h := range dmap.Height {
			pos.Y = h
			if dmap.GetValue(0, pos) {
				if isRollLiftable(&dmap, pos) {
					rollsLifted++
				}
			}
		}
	}
	fmt.Println("Rolls", rollsLifted)
}

func isRollLiftable(dmap *spatial.DiscreteMap2D[bool], pos spatial.DiscretePos2D) bool {
	rollNeihgbours := 0
	for _, p := range spatial.GetAdjacentPositions(pos) {
		if dmap.IsPositionInbounds(p) {
			if dmap.GetValue(0, p) {
				rollNeihgbours++
				if rollNeihgbours > 4 {
					// roll can be lifted
					return false
				}
			}
		}
	}
	return true
}

func createMap() spatial.DiscreteMap2D[bool] {
	lines := input.ReadLinesInFile("input.txt")

	dmap := spatial.NewDiscreteMap2DFromLines(1, lines, func(r rune) bool {
		switch r {
		case '@':
			return true
		default:
			return false
		}
	})
	return dmap
}
