package vector

import "math"

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

func (a *Vec2) Dot(b *Vec2) int {
	r := a.X*b.X + a.Y*b.Y
	return r
}

func (v *Vec2) Length() float64 {
	return math.Sqrt(float64(v.Dot(v)))
}
