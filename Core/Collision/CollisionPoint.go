package Collision

import "luigi.vanacore/go-snake/Core/Shapes"

type Point struct {
	Shapes.Point
}

func (p *Point) IsColliding(collision Collision) (error, bool) {

	return nil, false
}

func (p *Point) WillCollide(collision Collision, dx float64, dy float64) bool {
	pointTemp := p
	pointTemp.X += dx
	pointTemp.Y += dy
	_, coll := pointTemp.IsColliding(collision)
	return coll
}
