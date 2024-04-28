package main

import (
	"fmt"
	"os"

	"github.com/Flonka/advent-of-code/input"
	"github.com/Flonka/advent-of-code/spatial"
)

const size = 140

// ConnectingCell describes a pipe section, with data on where connections are positioned.
type Cell struct {
	connections [2]Direction
	textValue   rune
}

// HasConenction returns true if cell contains a connection to direction
func (c Cell) HasConnection(d Direction) bool {
	return c.connections[0] == d || c.connections[1] == d
}

func NewCell(r rune, c1 Direction, c2 Direction) Cell {
	c := Cell{}
	c.textValue = r
	c.connections[0] = c1
	c.connections[1] = c2
	return c
}

type Direction int

const (
	// Zero value for Direction (integer) is the first line
	Unitialized Direction = iota
	North
	East
	South
	West
)

const (
	CellValues = iota
	DistanceValues
)

func main() {

	s := input.OpenFileBuffered("input.txt")

	pipeMap := spatial.NewDiscreteMap2D[Cell](size, size, 1)

	y := 0
	startPos := spatial.DiscretePos2D{}
	pos := spatial.DiscretePos2D{}
	for s.Scan() {
		// Read all lines, create map
		// Find starting pos
		l := s.Text()
		cells := parseLine(l)

		// Assert correct size...
		if len(cells) != size {
			os.Exit(1)
		}

		pos.Y = y
		for x := 0; x < size; x++ {
			pos.X = x
			c := cells[x]
			if c.textValue == 'S' {
				startPos = pos
			}
			pipeMap.SetValue(CellValues, pos, c)
		}

		y++
	}

	fmt.Println("start:", startPos)
	findConnectedPipes(startPos, pipeMap)

	// Figure out where S connects toward. based on surrounding pipes
	// Go through map, starting at start , tracing positions, until reaching start again. Find position furthest
	// fmt.Println(pipeMap)
}

func findConnectedPipes(pos spatial.DiscretePos2D, dmap spatial.DiscreteMap2D[Cell]) []spatial.DiscretePos2D {

	connected := make([]spatial.DiscretePos2D, 0, 4)
	n := spatial.GetBorderPositions(pos)

	// x+1
	p := n[0]
	if dmap.IsPositionInbounds(p) {
		c := Cell(dmap.GetValue(CellValues, p))

		if c.HasConnection(West) {
			connected = append(connected, p)
		}
	}
	// x-1

	// y+1

	// y-1
	return connected
}

func parseLine(line string) []Cell {

	cells := make([]Cell, 0, len(line))

	for _, r := range line {

		c := cellFromRune(r)
		cells = append(cells, c)

	}

	return cells
}

func cellFromRune(r rune) Cell {
	switch r {
	case '|':
		return NewCell(r, North, South)
	case '-':
		return NewCell(r, West, East)
	case 'L':
		return NewCell(r, North, East)
	case 'J':
		return NewCell(r, West, North)
	case '7':
		return NewCell(r, West, South)
	case 'F':
		return NewCell(r, South, East)
	default:
		return Cell{textValue: r}
	}

}
