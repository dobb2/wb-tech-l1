package Point

import "math"

type Point struct {
	x float64
	y float64
}

func New(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt((p.x-p2.x)*(p.x-p2.x) + (p.y-p2.y)*(p.y-p2.y))
}
