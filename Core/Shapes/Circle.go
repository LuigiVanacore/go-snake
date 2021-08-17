package Shapes

import (
	"luigi.vanacore/go-snake/Core/Transform"
	"math"
)

type Circle struct {
	Transform.Position
	Radius float64
}

func NewCircle(x, y, radius float64) *Circle {
	return &Circle{Position: Transform.Position{X: x, Y: y}, Radius: radius}
}

func (c *Circle) GetPosition() (float64, float64) {
	return c.X, c.X
}

func (c *Circle) SetPosition(x float64, y float64) {
	c.X = x
	c.Y = y
}

func (c *Circle) Translate(x float64, y float64) {
	c.X += x
	c.Y += y
}

func (c *Circle) Equal(c2 *Circle) bool {
	return c.X == c2.X && c.Y == c2.Y && c.Radius == c2.Radius
}

func (c *Circle) Contains(x, y float64) bool {
	return math.Pow(x-c.X, 2)+math.Pow(y-c.Y, 2) < (c.Radius * c.Radius)
}
