package spatial

type DiscreteMap2D struct {
	Width  int
	Height int
	Data   [][]int
}

func NewDiscreteMap2D(width, height, dim int) DiscreteMap2D {

	d := make([][]int, dim)

	for i := 0; i < dim; i++ {
		d[i] = make([]int, width*height)
	}

	return DiscreteMap2D{
		Width:  width,
		Height: height,
		Data:   d,
	}
}

func (d *DiscreteMap2D) SetValue(dim int, pos DiscretePos2D, value int) {
	d.Data[dim][d.GetDataIndex(pos)] = value
}

func (d *DiscreteMap2D) GetDataIndex(pos DiscretePos2D) int {
	return pos.Y*d.Width + pos.X
}

type DiscretePos2D struct {
	X int
	Y int
}
