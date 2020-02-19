package main

import (
	"fmt"
	"math"
)

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	dis := Point.Distance
	fmt.Println(dis(p, q))
}

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
