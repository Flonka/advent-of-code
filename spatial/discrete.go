// Package spatial contain types about accessing and storing spatial data.
package spatial

import (
	"fmt"
	"log/slog"
)

// DiscreteMap2D can be used for storing 2d data with multiple dimensions in
// a integer coordinate system.
type DiscreteMap2D[T any] struct {
	Width  int
	Height int
	Data   [][]T
}

func (d DiscreteMap2D[T]) String() string {
	return fmt.Sprintf("{Width=%v,Height=%v}", d.Width, d.Height)
}

// SetValue sets a value of a position in the dimension
func (d *DiscreteMap2D[T]) SetValue(dim int, pos DiscretePos2D, value T) {
	d.Data[dim][d.GetDataIndex(pos)] = value
}

// GetValue the value of position in dimension.
// The positions x is width, y is height.
func (d *DiscreteMap2D[T]) GetValue(dim int, pos DiscretePos2D) T {
	return d.Data[dim][d.GetDataIndex(pos)]
}

// GetDataIndex gets the index to read data in the DiscreteMap2D
// position x is used for width, y for height
func (d *DiscreteMap2D[T]) GetDataIndex(pos DiscretePos2D) int {
	return pos.Y*d.Width + pos.X
}

func (d *DiscreteMap2D[T]) IsPositionInbounds(p DiscretePos2D) bool {
	if p.X < 0 || p.X >= d.Width {
		return false
	}
	if p.Y < 0 || p.Y >= d.Height {
		return false
	}
	return true
}

// DiscretePos2D is used to access values in the DiscreteMap2D.
// The x values are used for width, y for height.
type DiscretePos2D struct {
	X int
	Y int
}

var (
	N  = DiscretePos2D{0, 1}
	NE = DiscretePos2D{1, 1}
	E  = DiscretePos2D{1, 0}
	SE = DiscretePos2D{1, -1}
	S  = DiscretePos2D{0, -1}
	SW = DiscretePos2D{-1, -1}
	W  = DiscretePos2D{-1, 0}
	NW = DiscretePos2D{-1, 1}
)

// Add returns the added result as a new position
func (p DiscretePos2D) Add(p2 DiscretePos2D) DiscretePos2D {
	return DiscretePos2D{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

// NewDiscreteMap2D creates a new DiscreteMap2D with given width, height and dimension
func NewDiscreteMap2D[T any](width, height, dimension int) DiscreteMap2D[T] {
	d := make([][]T, dimension)

	for i := range dimension {
		d[i] = make([]T, width*height)
	}

	return DiscreteMap2D[T]{
		Width:  width,
		Height: height,
		Data:   d,
	}
}

// NewDiscreteMap2DFromLines creates DiscreteMap2D with width and height from the input lines
// it is assumed all linesare the same length
// first dimension is filled from lines in combination with the transformFunc data type.
func NewDiscreteMap2DFromLines[T any](dimension int, lines []string, transformFunc func(r rune, pos DiscretePos2D) T) DiscreteMap2D[T] {
	w := len(lines[0])
	h := len(lines)

	dmap := NewDiscreteMap2D[T](w, h, dimension)
	p := DiscretePos2D{}

	for y, l := range lines {
		p.Y = y
		for x, r := range l {
			p.X = x
			dmap.SetValue(0, p, transformFunc(r, p))
		}
	}

	slog.Debug("NewDiscreteMap2DFromLines", "map", dmap)
	return dmap
}

// GetBorderPositions returns the four bordering neighbour positions
// in order:
// x+1, x-1, y+1, y-1
// East, West, North, South
func GetBorderPositions(pos DiscretePos2D) []DiscretePos2D {
	return []DiscretePos2D{
		pos.Add(E),
		pos.Add(W),
		pos.Add(N),
		pos.Add(S),
	}
}

// GetAdjacentPositions returns the adjacent positions.
// in clockwise order starting from North
func GetAdjacentPositions(pos DiscretePos2D) []DiscretePos2D {
	return []DiscretePos2D{
		pos.Add(N),
		pos.Add(NE),
		pos.Add(E),
		pos.Add(SE),
		pos.Add(S),
		pos.Add(SW),
		pos.Add(W),
		pos.Add(NW),
	}
}
