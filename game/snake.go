package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	Up = iota
	Down
	Left
	Right
)

type SnakePart struct {
	BodyRect *ShapeNode
}

func (s *SnakePart) getNode() *ShapeNode {
	return s.BodyRect
}

type Snake struct {
	direction int
	size      int
	speed     int
	lives     int
	scores    int
	isDead    bool
	body      []*SnakePart
}

func (s *Snake) Init() {
	s.IncreaseBody()
}

func (s *Snake) IncreaseBody() {
	snakePart := SnakePart{&ShapeNode{Image: ebiten.NewImage(50, 50), Transform: ebiten.GeoM{}}}
	snakePart.getNode().getImage().Fill(color.RGBA{0xff, 0, 0, 0xff})
	s.body = append(s.body, &snakePart)
}

func (s *Snake) getBodyPart() *SnakePart {
	return s.body[0]
}
