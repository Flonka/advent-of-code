package main

import (
	"fmt"
	"os"

	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

const size = 140

type Pipe int

const (
	NorthSouth Pipe = iota
	EastWest
	NorthEast
	NorthWest
	SouthWest
	SouthEast
	Ground
	Start
	Unknown
)

func (p Pipe) String() string {
	return [...]string{"NS", "EW", "NE", "NW", "SW", "SE", "Ground", "Start", "Unknown"}[p]
}

const (
	PipeValues = iota
	DistanceValues
)

func main() {

	s := input.OpenFileBuffered("input.txt")

	pipeMap := spatial.NewDiscreteMap2D(size, size, 1)

	y := 0
	startPos := spatial.DiscretePos2D{}
	pos := spatial.DiscretePos2D{}
	for s.Scan() {
		// Read all lines, create map
		// Find starting pos
		l := s.Text()
		pipes := parseLineToPipes(l)

		// Assert correct size...
		if len(pipes) != size {
			os.Exit(1)
		}

		pos.Y = y
		for x := 0; x < size; x++ {
			pos.X = x
			p := pipes[x]
			if p == Start {
				startPos = pos
			}
			pipeMap.SetValue(PipeValues, pos, int(p))
		}

		y++
	}

	fmt.Println("start:", startPos)
	findConnectedPipes(startPos, pipeMap)

	// Figure out where S connects toward. based on surrounding pipes
	// Go through map, starting at start , tracing positions, until reaching start again. Find position furthest
	// fmt.Println(pipeMap)
}

func findConnectedPipes(pos spatial.DiscretePos2D, dmap spatial.DiscreteMap2D) []spatial.DiscretePos2D {

	connected := make([]spatial.DiscretePos2D, 0, 4)
	n := spatial.GetBorderPositions(pos)

	// x+1
	p := n[0]
	if inbounds(p, dmap) {
		pv := Pipe(dmap.GetValue(PipeValues, p))
		switch pv {

		case NorthWest, EastWest, SouthWest:
			connected = append(connected, p)

		}
	}
	// x-1

	// y+1

	// y-1
	return connected
}

func inbounds(p spatial.DiscretePos2D, dmap spatial.DiscreteMap2D) bool {

	if p.X < 0 || p.X > dmap.Width {
		return false
	}
	if p.Y < 0 || p.Y > dmap.Height {
		return false
	}
	return true
}

func parseLineToPipes(line string) []Pipe {

	pipes := make([]Pipe, 0, len(line))

	for _, r := range line {

		p := pipeFromRune(r)
		if p == Unknown {
			os.Exit(1)
		}
		pipes = append(pipes, p)

	}

	return pipes
}

func pipeFromRune(r rune) Pipe {
	switch r {
	case '|':
		return NorthSouth
	case '-':
		return EastWest
	case 'L':
		return NorthEast
	case 'J':
		return NorthWest
	case '7':
		return SouthWest
	case 'F':
		return SouthEast
	case '.':
		return Ground
	case 'S':
		return Start
	default:
		return Unknown
	}
}
