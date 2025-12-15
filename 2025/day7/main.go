package main

import (
	"errors"
	"log/slog"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/dutil"
	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

const SPLITTER int = 1

const TACHYION int = 2

func main() {
	cli.Default()
	dmap, start := createMap()

	splitCount := findSplits(&dmap, start)

	slog.Info("Part1", "splitCount", splitCount)
}

// findSplits returns the split count
func findSplits(dmap *spatial.DiscreteMap2D[int], startPos spatial.DiscretePos2D) int {
	pSet := dutil.NewSet[spatial.DiscretePos2D]()
	processed := dutil.NewSet[spatial.DiscretePos2D]()
	pSet.Add(startPos)
	processed.Add(startPos)

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
			p1 := splitPos.Add(spatial.E)
			p2 := splitPos.Add(spatial.W)

			// Dont add traces if already added
			// Dont add traces if the position is marked/traced by tachyons already

			v1 := dmap.GetValue(0, p1)
			v2 := dmap.GetValue(0, p2)

			if !processed.Contains(p1) && v1 != TACHYION {
				slog.Debug("Adding new trace", "pos", p1)
				pSet.Add(p1)
			}
			if !processed.Contains(p2) && v2 != TACHYION {
				slog.Debug("Adding new trace", "pos", p2)
				pSet.Add(p2)
			}

			processed.Add(p1)
			processed.Add(p2)

		}
		slog.Debug("Removing", "pos", pos)
		pSet.Remove(pos)
	}

	return count
}

// findSplitter traces North to find splitter
// and return the splitter position, or error if not found
// it also updates the map with tachyion traces
func findSplitter(dmap *spatial.DiscreteMap2D[int], p spatial.DiscretePos2D) (spatial.DiscretePos2D, error) {
	// trace the beam South(North in our map), per start position in the set of positions.
	// add to the set when a beam is split, and increment splitcounter

	dmap.SetValue(0, p, TACHYION)

	// Start adding to the start pos
	p = p.Add(spatial.N)
	// Loop until split found
	for {
		v := dmap.GetValue(0, p)
		switch v {
		case SPLITTER:
			return p, nil
		case TACHYION:
			return p, errors.New("ran into other trace")
		}

		dmap.SetValue(0, p, TACHYION)
		p = p.Add(spatial.N)
		if !dmap.IsPositionInbounds(p) {
			return p, errors.New("no splitter found")
		}
	}
}

func createMap() (spatial.DiscreteMap2D[int], spatial.DiscretePos2D) {
	lines := input.ReadLinesInFile("input.txt")

	var startPos spatial.DiscretePos2D

	dmap := spatial.NewDiscreteMap2DFromLines(1, lines, func(r rune, pos spatial.DiscretePos2D) int {
		if r == 'S' {
			startPos = pos
		}
		if r == '^' {
			return SPLITTER
		}
		return 0
	})
	return dmap, startPos
}
