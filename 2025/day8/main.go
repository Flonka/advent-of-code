package main

import (
	"log/slog"
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
}

func readInput() []mathutils.Vec3[int] {
	boxes := make([]mathutils.Vec3[int], 0, 500)
	input.HandleInputLines(func(lineCount, lineIndex int, line string) {
		c := input.StringsToInts(strings.Split(line, ","))
		boxes = append(boxes, mathutils.Vec3[int]{X: c[0], Y: c[1], Z: c[2]})
	})

	return boxes
}
