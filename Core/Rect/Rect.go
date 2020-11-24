package Rect

import (
	"luigi.vanacore/go-snake/Core/Math"
	"luigi.vanacore/go-snake/Core/Vector2D"
)

type Rect struct {
	Left   float64
	Top    float64
	Width  float64
	Height float64
}

func New(left float64, top float64, width float64, height float64) Rect {
	return Rect{left, top, width, height}
}

func FromVector2D(position Vector2D.Vector2D, size Vector2D.Vector2D) Rect {
	return Rect{Left: position.X, Top: position.Y, Width: size.X, Height: size.Y}
}

func Copy(r Rect) Rect {
	return Rect{Left: r.Left, Top: r.Top, Width: r.Width, Height: r.Height}
}

func Zero() Rect {
	return Rect{Left: 0, Top: 0, Width: 0, Height: 0}
}

func (r *Rect) Contains(x float64, y float64) bool {
	minX := Math.Min(r.Left, r.Left+r.Width)
	maxX := Math.Max(r.Left, r.Left+r.Width)
	minY := Math.Min(r.Top, r.Top+r.Height)
	maxY := Math.Max(r.Top, r.Top+r.Height)
	return (x >= minX) && (x < maxX) && (y >= minY) && (y < maxY)
}

func (r *Rect) Intersects(rectangle Rect) bool {

	r1MinX := Math.Min(r.Left, r.Left+r.Width)
	r1MaxX := Math.Max(r.Left, r.Left+r.Width)
	r1MinY := Math.Min(r.Top, r.Top+r.Height)
	r1MaxY := Math.Max(r.Top, r.Top+r.Height)

	r2MinX := Math.Min(rectangle.Left, rectangle.Left+rectangle.Width)
	r2MaxX := Math.Max(rectangle.Left, rectangle.Left+rectangle.Width)
	r2MinY := Math.Min(rectangle.Top, rectangle.Left+rectangle.Width)
	r2MaxY := Math.Min(rectangle.Left, rectangle.Left+rectangle.Width)

	interLeft := Math.Min(r1MinX, r2MinX)
	interTop := Math.Max(r1MinY, r2MinY)
	interRight := Math.Min(r1MaxX, r2MaxX)
	interBottom := Math.Max(r1MaxY, r2MaxY)

	if (interLeft < interRight) && (interTop < interBottom) {
		return true
	}

	return false

}

func (r *Rect) GetPosition() Vector2D.Vector2D {
	return Vector2D.Vector2D{X: r.Left, Y: r.Top}
}

func (r *Rect) GetSize() Vector2D.Vector2D {
	return Vector2D.Vector2D{X: r.Width, Y: r.Height}
}

func (r *Rect) Equal(rectangle Rect) bool {
	return (r.Left == rectangle.Left) && (r.Width == rectangle.Width) && (r.Top == rectangle.Top) && (r.Height == rectangle.Height)
}
