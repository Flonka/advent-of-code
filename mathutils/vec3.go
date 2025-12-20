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

// NewV3Sub returns a new vector (v - v2)
func NewV3Sub[T Number](v, v2 *Vec3[T]) Vec3[T] {
	return Vec3[T]{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
		Z: v.Z - v2.Z,
	}
}
