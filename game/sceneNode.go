package game

import "github.com/hajimehoshi/ebiten/v2"

type ShapeNode struct {
	Image     *ebiten.Image
	Transform ebiten.GeoM
}

func (s *ShapeNode) GetImage() *ebiten.Image {
	return s.Image
}

func (s *ShapeNode) SetTransform(t *ebiten.GeoM) {
	s.Transform = *t
}

func (s *ShapeNode) GetTransform() *ebiten.GeoM {
	return &s.Transform
}
