package mathutils

import "math"

type Number interface {
	~int | float64 | float32
}

// Vec3 generic vector in three dimensions.
type Vec3[T Number] struct {
	X T
	Y T
	Z T
}

func (v *Vec3[T]) Length() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}
