package spatial

type DiscreteMap2D[T any] struct {
	Width  int
	Height int
	Data   [][]T
}

// SetValue sets a value of a position in the dimension
func (d *DiscreteMap2D[T]) SetValue(dim int, pos DiscretePos2D, value T) {
	d.Data[dim][d.GetDataIndex(pos)] = value
}

// Get the value of position in dimension
func (d *DiscreteMap2D[T]) GetValue(dim int, pos DiscretePos2D) T {
	return d.Data[dim][d.GetDataIndex(pos)]
}

func (d *DiscreteMap2D[T]) GetDataIndex(pos DiscretePos2D) int {
	return pos.Y*d.Width + pos.X
}

func (d *DiscreteMap2D[T]) IsPositionInbounds(p DiscretePos2D) bool {

	if p.X < 0 || p.X > d.Width {
		return false
	}
	if p.Y < 0 || p.Y > d.Height {
		return false
	}
	return true
}

type DiscretePos2D struct {
	X int
	Y int
}

func (p DiscretePos2D) Add(p2 DiscretePos2D) DiscretePos2D {
	return DiscretePos2D{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

// NewDiscreteMap2D creates a new DiscreteMap2D with given width, height and dimension
func NewDiscreteMap2D[T any](width, height, dimension int) DiscreteMap2D[T] {

	d := make([][]T, dimension)

	for i := 0; i < dimension; i++ {
		d[i] = make([]T, width*height)
	}

	return DiscreteMap2D[T]{
		Width:  width,
		Height: height,
		Data:   d,
	}
}

// GetBorderPositions returns the four bordering neighbour positions
// in order:
// x+1, x-1, y+1, y-1
func GetBorderPositions(pos DiscretePos2D) []DiscretePos2D {

	return []DiscretePos2D{
		{X: pos.X + 1, Y: pos.Y},
		{X: pos.X - 1, Y: pos.Y},
		{X: pos.X, Y: pos.Y + 1},
		{X: pos.X, Y: pos.Y - 1},
	}

}
