package spatial

type DiscreteMap2D struct {
	Width  int
	Height int
	Data   [][]int
}

func (d *DiscreteMap2D) SetValue(dim int, pos DiscretePos2D, value int) {
	d.Data[dim][d.GetDataIndex(pos)] = value
}

func (d *DiscreteMap2D) GetValue(dim int, pos DiscretePos2D) int {
	return d.Data[dim][d.GetDataIndex(pos)]
}

func (d *DiscreteMap2D) GetDataIndex(pos DiscretePos2D) int {
	return pos.Y*d.Width + pos.X
}

type DiscretePos2D struct {
	X int
	Y int
}

// NewDiscreteMap2D creates a new DiscreteMap2D with given width, height and dimension
func NewDiscreteMap2D(width, height, dimension int) DiscreteMap2D {

	d := make([][]int, dimension)

	for i := 0; i < dimension; i++ {
		d[i] = make([]int, width*height)
	}

	return DiscreteMap2D{
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
