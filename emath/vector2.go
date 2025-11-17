package emath

import "github.com/chewxy/math32"

type Vector2 struct {
	X, Y float32
}

func Lerp(start float32, end float32, t float32) float32 {
	return start + (end-start)*t
}

func (v Vector2) DistanceTo(v2 Vector2) float32 {
	return math32.Sqrt((v2.X-v.X)*(v2.X-v.X) + (v2.Y-v.Y)*(v2.Y-v.Y))
}

func (v Vector2) Sub(v2 Vector2) Vector2 {
	return NewVector2(v.X-v2.X, v.Y-v2.Y)
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	return NewVector2(v.X+v2.X, v.Y+v2.Y)
}

func (v Vector2) Mul(n float32) Vector2 {
	return NewVector2(v.X*n, v.Y*n)
}

func (v Vector2) Invert() Vector2 {
	return NewVector2(-v.X, -v.Y)
}

func (v Vector2) Lerp(v2 Vector2, t float32) Vector2 {
	return v.Add(v2.Sub(v)).Mul(t)
}

func NewVector2(x float32, y float32) Vector2 {
	return Vector2{x, y}
}
