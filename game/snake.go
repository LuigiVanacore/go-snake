package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Direction int

const (
	Up = iota
	Down
	Left
	Right
	None
)

type SnakePart struct {
	BodyRect *ShapeNode
}

func (s *SnakePart) GetNode() *ShapeNode {
	return s.BodyRect
}

type Snake struct {
	direction Direction
	size      int
	speed     int
	lives     int
	score     int
	isDead    bool
	body      []*SnakePart
}

func (s *Snake) Init() {
	transform := ebiten.GeoM{}
	s.IncreaseBody(&transform)
}

func (s *Snake) IncreaseBody(m *ebiten.GeoM) {
	snakePart := SnakePart{&ShapeNode{Image: ebiten.NewImage(50, 50), Transform: ebiten.GeoM{}}}
	snakePart.GetNode().GetImage().Fill(color.RGBA{0xff, 0, 0, 0xff})
	snakePart.GetNode().SetTransform(m)
	s.body = append(s.body, &snakePart)
}

func (s *Snake) GetBodyPart() *SnakePart {
	return s.body[0]
}

func (s *Snake) SetDirection(direction Direction) {
	s.direction = direction
}

func (s *Snake) GetDirection() Direction {
	return s.direction
}

func (s *Snake) GetSpeed() int {
	return s.speed
}

func (s *Snake) GetLives() int {
	return s.lives
}

func (s *Snake) GetScore() int {
	return s.score
}

func (s *Snake) IncreaseScore() {
	s.score++
}

func (s *Snake) HasLost() bool {
	return s.isDead
}

func (s *Snake) Reset() {
	s.body = nil
	transform := ebiten.GeoM{}
	s.IncreaseBody(&transform)
	s.IncreaseBody(&transform)
	s.IncreaseBody(&transform)
	s.SetDirection(None)
	s.speed = 15
	s.lives = 3
	s.score = 0
	s.isDead = false
}
