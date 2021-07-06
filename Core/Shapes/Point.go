package Shapes

import "luigi.vanacore/go-snake/Core/Collision"

type Point struct {
	X, Y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) GetPosition() (float64, float64) {
	return p.X, p.Y
}

func (p *Point) SetPosition(x float64, y float64) {
	p.X = x
	p.Y = y
}

func (p *Point) Translate(x float64, y float64) {
	p.X += x
	p.Y += y
}

func (p *Point) Equal(p2 *Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p *Point) IsColliding(collision Collision.Collision) bool {

	return false
}

func (p *Point) WillCollide(collision Collision.Collision, dx float64, dy float64) bool {
	pointTemp := p
	pointTemp.X += dx
	pointTemp.Y += dy
	return pointTemp.IsColliding(collision)
}
