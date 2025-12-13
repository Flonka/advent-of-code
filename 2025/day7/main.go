package main

import (
	"log/slog"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

func main() {
	cli.Default()
	dmap, start := createMap()
	slog.Info("Data read", "pos", start, "map", dmap)
}

func createMap() (spatial.DiscreteMap2D[rune], spatial.DiscretePos2D) {
	lines := input.ReadLinesInFile("input.txt")

	var startPos spatial.DiscretePos2D

	dmap := spatial.NewDiscreteMap2DFromLines(1, lines, func(r rune, pos spatial.DiscretePos2D) rune {
		if r == 'S' {
			startPos = pos
		}
		return r
	})
	return dmap, startPos
}
