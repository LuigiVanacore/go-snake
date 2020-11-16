package Shapes

import "luigi.vanacore/go-snake/Core/Vector2D"

type RectangleShape struct {
	size Vector2D.Vector2D
	Shape
}

func New(size Vector2D.Vector2D) RectangleShape {
	return RectangleShape{size: size}
}

func (rect *RectangleShape) SetSize(size Vector2D.Vector2D) {
	rect.size = size
}

func (rect *RectangleShape) GetSize() Vector2D.Vector2D {
	return rect.size
}

func (rect *RectangleShape) GetPointCount() uint32 {
	//TODO
	return 0
}

func (rect *RectangleShape) GetPoint(index uint) Vector2D.Vector2D {
	//TODO
	return Vector2D.Zero()
}
