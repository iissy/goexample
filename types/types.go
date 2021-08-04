package types

import "math"

type Point struct {
	Dian
}

type Dian struct {
	X, Y float64
}

func (p Dian) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
