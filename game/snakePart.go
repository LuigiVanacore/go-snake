package game

import (
	"luigi.vanacore/go-snake/Core/Transform"
)

type SnakePart struct {
	position Transform.Position
}

func NewSnakePart(position Transform.Position) *SnakePart {
	return &SnakePart{position: position}
}

func (s *SnakePart) GetPosition() Transform.Position {
	return s.position
}

func (s *SnakePart) SetPosition(x, y float64) {
	s.position.X = x
	s.position.Y = y
}

func (s *SnakePart) Move(direction int) {
	switch direction {
	case Left:
		s.position.X--
	case Right:
		s.position.X++
	case Down:
		s.position.Y++
	case Up:
		s.position.Y--
	}
}

//func (s *SnakePart) Draw(screen *ebiten.Image) {
//	op := &ebiten.DrawImageOptions{}
//	op.GeoM.Translate(s.Position.X, s.Position.Y)
//	screen.DrawImage(s.ShapeNode.GetImage(), op)
//}
