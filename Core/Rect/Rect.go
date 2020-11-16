package Rect

import "luigi.vanacore/go-snake/Core/Vector2D"

type Rect struct {
	Left   float32
	Top    float32
	Width  float32
	Height float32
}

func New(left float32, top float32, width float32, height float32) Rect {
	return Rect{left, top, width, height}
}

func FromVector2D(position Vector2D.Vector2D, size Vector2D.Vector2D) Rect {
	return Rect{Left: position.X, Top: position.Y, Width: size.X, Height: size.Y}
}

func Copy(r Rect) Rect {
	return Rect{Left: r.Left, Top: r.Top, Width: r.Width, Height: r.Height}
}

func (r *Rect) Contains(x float32, y float32) bool {
	//TODO
	return false
}

func (r *Rect) Intersects(rectangle Rect, intersection Rect) bool {
	//TODO
	return false
}

func (r *Rect) GetPosition() Vector2D.Vector2D {
	//TODO
	return Vector2D.Zero()
}

func (r *Rect) GetSize() bool {
	//TODO
	return false
}

func (r *Rect) Equal(rectangle Rect) bool {
	//TODO
	return false
}
