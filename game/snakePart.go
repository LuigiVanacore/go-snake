package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"luigi.vanacore/go-snake/Core/Vector2D"
)

type SnakePart struct {
	ShapeNode *ShapeNode
	Position  Vector2D.Vector2D
}

func New(body *ShapeNode, position Vector2D.Vector2D) *SnakePart {
	return &SnakePart{ShapeNode: body, Position: position}
}

func (s *SnakePart) GetNode() *ShapeNode {
	return s.ShapeNode
}

func (s *SnakePart) GetPosition() Vector2D.Vector2D {
	return s.Position
}

func (s *SnakePart) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.Position.X, s.Position.Y)
	screen.DrawImage(s.ShapeNode.GetImage(), op)
}
