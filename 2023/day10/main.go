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
	// S type cell connects to every direction
	if c.textValue == 'S' {
		return true
	}
	return c.connections[0] == d || c.connections[1] == d
}

func (c Cell) GetConnectingPositions() []spatial.DiscretePos2D {

	p1 := c.connections[0].ToPosition()
	p2 := c.connections[1].ToPosition()
	return []spatial.DiscretePos2D{p1, p2}

}

func NewCell(r rune, c1 Direction, c2 Direction) Cell {
	c := Cell{}
	c.textValue = r
	c.connections[0] = c1
	c.connections[1] = c2
	return c
}

type Direction int

func (d Direction) ToPosition() spatial.DiscretePos2D {
	switch d {
	case North:
		return spatial.DiscretePos2D{X: 0, Y: -1}
	case South:
		return spatial.DiscretePos2D{X: 0, Y: 1}
	case East:
		return spatial.DiscretePos2D{X: 1, Y: 0}
	case West:
		return spatial.DiscretePos2D{X: -1, Y: 0}
	default:
		panic("not implemented")
	}
}

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

	maze := spatial.NewDiscreteMap2D[Cell](size, size, 1)

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
			maze.SetValue(CellValues, pos, c)
		}

		y++
	}

	fmt.Println("start:", startPos)

	// Pick one to start from, go until we reach start pos
	traverseLoop(maze, startPos)

}

func traverseLoop(maze spatial.DiscreteMap2D[Cell], start spatial.DiscretePos2D) {
	// Figure out where S connects toward. based on surrounding pipes
	startConns := findConnectedPipes(start, maze)

	stepCount := 0

	// Pick one path to start going
	pos := startConns[0]
	lastPos := start
	for {

		fmt.Println("Current pos", pos)
		cell := maze.GetValue(CellValues, pos)
		next := stepFurther(cell, pos, lastPos)
		lastPos = pos
		pos = next
		stepCount++
		fmt.Println(stepCount)
		// Return if we have reached start
		if pos == start {
			fmt.Println("Found start again", pos, start)
			break
		}
	}

	fmt.Println("steps", stepCount)
}

// Return next pos
func stepFurther(cell Cell, pos spatial.DiscretePos2D, last spatial.DiscretePos2D) spatial.DiscretePos2D {

	// Get connecting positions from cell, return next one not being the last one.
	for _, p := range cell.GetConnectingPositions() {
		px := pos.Add(p)
		if px != last {
			return px
		}
	}

	panic("not found new pos")

}

// findConnectedPipes returns a slice of positions which have cells connecting to the given position
func findConnectedPipes(pos spatial.DiscretePos2D, dmap spatial.DiscreteMap2D[Cell]) []spatial.DiscretePos2D {

	connected := make([]spatial.DiscretePos2D, 0, 4)
	n := spatial.GetBorderPositions(pos)

	// x+1
	p := n[0]
	if isConnected(dmap, p, West) {
		connected = append(connected, p)
	}
	// x-1
	p = n[1]
	if isConnected(dmap, p, East) {
		connected = append(connected, p)
	}
	// y+1
	p = n[2]
	if isConnected(dmap, p, North) {
		connected = append(connected, p)
	}

	// y-1
	p = n[3]
	if isConnected(dmap, p, South) {
		connected = append(connected, p)
	}
	return connected
}

func isConnected(dmap spatial.DiscreteMap2D[Cell], p spatial.DiscretePos2D, d Direction) bool {

	if dmap.IsPositionInbounds(p) {
		c := Cell(dmap.GetValue(CellValues, p))
		if c.HasConnection(d) {
			return true
		}
	}
	return false
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
