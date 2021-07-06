package Shapes

import (
	"luigi.vanacore/go-snake/Core/Collision"
	"luigi.vanacore/go-snake/Core/Utils"
)

type Rect struct {
	Point
	Width  float64
	Height float64
}

func NewRect(x float64, y float64, width float64, height float64) *Rect {
	return &Rect{Point{x, y}, width, height}
}

func (r *Rect) GetPosition() (float64, float64) {
	return r.X, r.Y
}

func (r *Rect) SetPosition(x float64, y float64) {
	r.X = x
	r.Y = y
}

func (r *Rect) Translate(x float64, y float64) {
	r.X += x
	r.Y += y
}

func (r *Rect) GetSize() float64 {
	return r.Width + r.Height
}

func (r *Rect) GetCenter() (float64, float64) {
	return r.X + r.Width/2, r.Y + r.Height/2
}
func (r *Rect) Equal(rectangle Rect) bool {
	return (r.X == rectangle.X) && (r.Width == rectangle.Width) && (r.Y == rectangle.Y) && (r.Height == rectangle.Height)
}

func (r *Rect) IsColliding(collision Collision.Collision) bool {
	switch c := collision.(type) {
	case *Point:
		return r.Contains(c.X, c.Y)
	case *Rect:
		return r.IntersectRect(c)
	case *Circle:
	default:
		return false
	}
	return false
}

func (r *Rect) WillCollide(collision Collision.Collision, dx float64, dy float64) bool {
	rectTemp := r
	rectTemp.X += dx
	rectTemp.Y += dy
	return rectTemp.IsColliding(collision)
}

func (r *Rect) Contains(x float64, y float64) bool {
	minX := Utils.Min(r.X, r.X+r.Width)
	maxX := Utils.Max(r.X, r.X+r.Width)
	minY := Utils.Min(r.Y, r.Y+r.Height)
	maxY := Utils.Max(r.Y, r.Y+r.Height)
	return (x >= minX) && (x < maxX) && (y >= minY) && (y < maxY)
}

func (r *Rect) IntersectRect(rectangle *Rect) bool {

	r1MinX := Utils.Min(r.X, r.X+r.Width)
	r1MaxX := Utils.Max(r.X, r.X+r.Width)
	r1MinY := Utils.Min(r.Y, r.Y+r.Height)
	r1MaxY := Utils.Max(r.Y, r.Y+r.Height)

	r2MinX := Utils.Min(rectangle.X, rectangle.X+rectangle.Width)
	r2MaxX := Utils.Max(rectangle.X, rectangle.X+rectangle.Width)
	r2MinY := Utils.Min(rectangle.Y, rectangle.X+rectangle.Width)
	r2MaxY := Utils.Min(rectangle.X, rectangle.X+rectangle.Width)

	interLeft := Utils.Min(r1MinX, r2MinX)
	interTop := Utils.Max(r1MinY, r2MinY)
	interRight := Utils.Min(r1MaxX, r2MaxX)
	interBottom := Utils.Max(r1MaxY, r2MaxY)

	if (interLeft < interRight) && (interTop < interBottom) {
		return true
	}

	return false

}
