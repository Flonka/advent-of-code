package day12

import (
	"github.com/Flonka/advent-of-code/spatial"
)

type ClimbMap struct {
	MapData spatial.DiscreteMap2D

	Start spatial.DiscretePos2D
	End   spatial.DiscretePos2D
}
