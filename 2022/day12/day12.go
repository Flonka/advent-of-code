package day12

import (
	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

type ClimbMap struct {
	MapData spatial.DiscreteMap2D

	Start spatial.DiscretePos2D
	End   spatial.DiscretePos2D
}

func ReadMapData(fname string) ClimbMap {

	lines := input.ReadLinesInFile(fname)

	var w, h int

	// assume all lines have equal length
	w = len(lines[0])
	h = len(lines)

	d := ClimbMap{
		MapData: spatial.NewDiscreteMap2D(w, h, 2),
	}

	start := int('S')
	end := int('E')

	var p spatial.DiscretePos2D
	for i := 0; i < h; i++ {
		p.Y = i
		l := lines[i]
		// Convert line strings to integers
		for j := 0; j < w; j++ {
			p.X = j
			v := int(l[j])
			// Handle Start&End
			switch v {
			case start:
				d.Start.X = p.X
				d.Start.Y = p.Y
			case end:
				d.End.X = p.X
				d.End.Y = p.Y
			}
			d.MapData.SetValue(0, p, v)
		}

	}

	return d
}
