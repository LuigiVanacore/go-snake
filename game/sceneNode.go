package game

import "github.com/hajimehoshi/ebiten/v2"

type ShapeNode struct {
	Image     *ebiten.Image
	Transform ebiten.GeoM
}

func (s *ShapeNode) getImage() *ebiten.Image {
	return s.Image
}
