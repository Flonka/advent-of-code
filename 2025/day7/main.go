package main

import (
	"log/slog"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/dutil"
	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

func main() {
	cli.Default()
	dmap, start := createMap()
	splitCount := 0

	pSet := dutil.NewSet[spatial.DiscretePos2D]()
	pSet.Add(start)

	slog.Info("start", "pos", start)

	// trace the beam South(North in our map), per start position in the set of positions.
	// add to the set when a beam is split, and increment splitcounter
	if dmap.GetValue(0, start.Add(spatial.N)) {
		slog.Info("asdf")
	}
	slog.Info("Part1", "splitCount", splitCount)
}

func createMap() (spatial.DiscreteMap2D[bool], spatial.DiscretePos2D) {
	lines := input.ReadLinesInFile("input.txt")

	var startPos spatial.DiscretePos2D

	dmap := spatial.NewDiscreteMap2DFromLines(1, lines, func(r rune, pos spatial.DiscretePos2D) bool {
		if r == 'S' {
			startPos = pos
		}
		if r == '^' {
			return true
		}
		return false
	})
	return dmap, startPos
}
