package mathutils

import "math"

// Vec2 of integer components.
type Vec2 struct {
	X int
	Y int
}

func (v *Vec2) Add(v2 *Vec2) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vec2) Subtract(v2 *Vec2) {
	v.X -= v2.X
	v.Y -= v2.Y
}

func (v *Vec2) Dot(b *Vec2) int {
	r := v.X*b.X + v.Y*b.Y
	return r
}

func (v *Vec2) Length() float64 {
	return math.Sqrt(float64(v.Dot(v)))
}
