package Collision

import (
	"fmt"
	"luigi.vanacore/go-snake/Core/Shapes"
	"luigi.vanacore/go-snake/Core/Utils"
)

type Circle struct {
	Shapes.Circle
}

func (c *Circle) IsColliding(collision Collision) (error, bool) {
	switch coll := collision.(type) {
	case *Point:
		return nil, c.Contains(c.X, c.Y)
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
		return nil, Utils.Distance(c.X, c.Y, closestX, closestY) <= c.Radius
	case *Circle:
		return nil, Utils.Distance(c.X, c.Y, coll.X, coll.Y) <= c.Radius+coll.Radius
	default:
		return fmt.Errorf("collision object provided not valid %d", collision), false
	}
}

func (c *Circle) WillCollide(collision Collision, dx float64, dy float64) bool {
	circleTemp := c
	circleTemp.X += dx
	circleTemp.Y += dy
	_, coll := circleTemp.IsColliding(collision)
	return coll
}
