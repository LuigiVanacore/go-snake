package Collision

import (
	"fmt"
	"luigi.vanacore/go-snake/Core/Shapes"
)

type Rect struct {
	Shapes.Rect
}

func (r *Rect) IsColliding(collision Collision) (error, bool) {
	switch c := collision.(type) {
	case *Point:
		return nil, r.Contains(c.X, c.Y)
	case *Rect:
		return nil, r.IntersectRect(&c.Rect)
	case *Circle:
	}
	return fmt.Errorf("collision object provided not valid", collision), false
}

func (r *Rect) WillCollide(collision Collision, dx float64, dy float64) bool {
	rectTemp := r
	rectTemp.X += dx
	rectTemp.Y += dy
	_, coll := rectTemp.IsColliding(collision)
	return coll
}
