package Shapes

import (
	"luigi.vanacore/go-snake/Core/Transform"
)

type Point struct {
	Transform.Position
}

func NewPoint(x float64, y float64) *Point {
	return &Point{Position: Transform.Position{X: x, Y: y}}
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
