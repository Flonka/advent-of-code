package main

import (
	"log/slog"
	"math"
	"slices"
	"strings"

	"github.com/Flonka/advent-of-code/cli"
	"github.com/Flonka/advent-of-code/dutil"
	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/mathutils"
)

func main() {
	cli.Default()

	boxes := readInput()

	slog.Info("boxes", "length", len(boxes))
	// connect together the 1000 pairs of junction boxes which are closest together. Afterward, what do you get if you multiply together the sizes of the three largest circuits?

	// Create the 1000 pairs
	pairs := make([]pair, 0, 1000)
	for i, b := range boxes {
		// Find the closest box for every box
		cb, l := findClosest(b, i, boxes)
		p := pair{i, cb, l}
		pairs = append(pairs, p)
		slog.Debug("Pair added", "pair", p)
	}

	slog.Info("pair length", "pairs", len(pairs))

	// a circuit can be seen as a set of indices
	circuits := make([]dutil.Set[int], 0, 10)

	// Add every pair into one circuit.
	for _, p := range pairs {
		// check if any of the pair's boxes are in a circuit.
		cI := findPairInCircuit(p, circuits)
		if cI != -1 {
			// add to existing circuit
			circuits[cI].Add(p.a)
			circuits[cI].Add(p.b)
		} else {
			// Otherwise, create a new circuit.
			newCircuit := *dutil.NewSet[int]()
			newCircuit.Add(p.a)
			newCircuit.Add(p.b)
			circuits = append(circuits, newCircuit)
		}
	}

	// sort.Slice(circuits, func(i int, j int) bool {
	// 	return circuits[i].Size() < circuits[j].Size()
	// })
	// Sort descending
	slices.SortFunc(circuits, func(a, b dutil.Set[int]) int {
		if a.Size() < b.Size() {
			return 1
		}
		return -1
	})

	slog.Info("Circuits", "len", len(circuits))

	p1 := circuits[0].Size() * circuits[1].Size() * circuits[2].Size()
	slog.Info("Part1", "answer", p1)
}

// findPairInCircuit returns circuit index if any of the pairs
// indices already are in some circuit
// -1 if not
func findPairInCircuit(p pair, circuits []dutil.Set[int]) int {
	for i, c := range circuits {
		if c.Contains(p.a) || c.Contains(p.b) {
			return i
		}
	}

	return -1
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
	// Index of a
	a int
	// Index of b
	b int
	// distance between the boxes of indices a and b.
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
