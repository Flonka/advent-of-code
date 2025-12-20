package main

import (
	"log/slog"
	"math"
	"strings"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/mathutils"
)

func main() {
	cli.Default()

	boxes := readInput()

	slog.Info("boxes", "length", len(boxes))
	// connect together the 1000 pairs of junction boxes which are closest together. Afterward, what do you get if you multiply together the sizes of the three largest circuits?

	pairs := make([]pair, 0, 1000)
	// Find the closest box for every box
	for i, b := range boxes {
		cb, l := findClosest(b, i, boxes)
		p := pair{i, cb, l}
		pairs = append(pairs, p)
		slog.Debug("Pair added", "pair", p)
	}

	slog.Info("pair length", "pairs", len(pairs))
}

// findClosest returns index for closest vector in the collection ,skipping vIndex
func findClosest(v mathutils.Vec3[int], vIndex int, collection []mathutils.Vec3[int]) (int, float64) {
	min := math.Inf(1)
	minIdx := -1
	for i, b := range collection {

		if i == vIndex {
			continue
		}
		dv := mathutils.NewV3Sub(&b, &v)
		d := dv.Length()
		if d < min {
			minIdx = i
			min = d
		}
	}
	return minIdx, min
}

// pair connect two indices
type pair struct {
	a        int
	b        int
	distance float64
}

func readInput() []mathutils.Vec3[int] {
	boxes := make([]mathutils.Vec3[int], 0, 500)
	input.HandleInputLines(func(lineCount, lineIndex int, line string) {
		c := input.StringsToInts(strings.Split(line, ","))
		boxes = append(boxes, mathutils.Vec3[int]{X: c[0], Y: c[1], Z: c[2]})
	})

	return boxes
}
