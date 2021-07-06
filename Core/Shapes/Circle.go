package Shapes

import (
	"luigi.vanacore/go-snake/Core/Collision"
	"luigi.vanacore/go-snake/Core/Utils"
	"math"
)

type Circle struct {
	Point
	Radius float64
}

func NewCircle(x, y, radius float64) *Circle {
	return &Circle{Point{x, y}, radius}
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

func (c *Circle) IsColliding(collision Collision.Collision) (error, bool) {
	switch coll := collision.(type) {
	case *Point:
		return _, c.Contains(c.X, c.Y)
	case *Rect:
		closestX := c.X
		closestY := c.Y

		if c.X < coll.X {
			closestX = coll.X
		} else if c.X > coll.X+coll.Width {
			closestX = coll.X + coll.Width
		}

		if c.Y < coll.Y {
			closestY = coll.Y
		} else if c.Y > coll.Y+coll.Height {
			closestY = coll.Y + coll.Height
		}
		return Utils.Distance(c.X, c.Y, closestX, closestY) <= c.Radius
	case *Circle:
		return Utils.Distance(c.X, c.Y, coll.X, coll.Y) <= c.Radius+coll.Radius
	default:
		return false
	}
}

func (c *Circle) Contains(x, y float64) bool {
	return math.Pow(x-c.X, 2)+math.Pow(y-c.Y, 2) < (c.Radius * c.Radius)
}

func (c *Circle) WillCollide(collision Collision.Collision, dx float64, dy float64) bool {
	circleTemp := c
	circleTemp.X += dx
	circleTemp.Y += dy
	return circleTemp.IsColliding(collision)
}
