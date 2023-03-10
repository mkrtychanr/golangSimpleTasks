package point

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p *Point) SetX(x float64) {
	p.x = x
}

func (p *Point) SetY(y float64) {
	p.y = y
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (firstPoint *Point) GetDistance(secondPoint *Point) float64 {
	x := math.Pow((secondPoint.x - firstPoint.x), 2)
	y := math.Pow((secondPoint.y - firstPoint.y), 2)
	return math.Sqrt(x + y)
}
