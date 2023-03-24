package spatial

type DiscreteMap2D struct {
	Width  int
	Height int
	data   [][]int
}

func NewDiscreteMap2D(width, height, dim int) DiscreteMap2D {

	d := make([][]int, dim)

	for i := 0; i < dim; i++ {
		d[i] = make([]int, width*height)
	}

	return DiscreteMap2D{
		Width:  width,
		Height: height,
		data:   d,
	}
}

type DiscretePos2D struct {
	X int
	Y int
}
