package main

import (
	"errors"
	"log/slog"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/dutil"
	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

func main() {
	cli.Default()
	dmap, start := createMap()

	splitCount := findSplits(&dmap, start)

	slog.Info("Part1", "splitCount", splitCount)
}

// findSplits returns the split count
func findSplits(dmap *spatial.DiscreteMap2D[bool], startPos spatial.DiscretePos2D) int {
	pSet := dutil.NewSet[spatial.DiscretePos2D]()
	pSet.Add(startPos)

	count := 0

	for pSet.Size() > 0 {
		slog.Debug("Set size", "size", pSet.Size())
		pos := pSet.List()[0]
		splitPos, err := findSplitter(dmap, pos)
		if err != nil {
			// No spliter found
			slog.Debug("trace OB", "pos", splitPos)
		} else {
			// splitter found, add new traces west and east
			count++
			slog.Debug("Adding new traces", "pos", splitPos)
			pSet.Add(splitPos.Add(spatial.E))
			pSet.Add(splitPos.Add(spatial.W))
		}
		slog.Debug("Removing", "pos", pos)
		pSet.Remove(pos)
	}

	return count
}

// findSplitter traces North to find splitter
// and return the splitter position, or error if not found
func findSplitter(dmap *spatial.DiscreteMap2D[bool], p spatial.DiscretePos2D) (spatial.DiscretePos2D, error) {
	// trace the beam South(North in our map), per start position in the set of positions.
	// add to the set when a beam is split, and increment splitcounter

	// Start adding to the start pos
	p = p.Add(spatial.N)
	for !dmap.GetValue(0, p) {
		p = p.Add(spatial.N)
		if !dmap.IsPositionInbounds(p) {
			return p, errors.New("no splitter found")
		}
	}
	// Should be a splitter pos now
	return p, nil
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
