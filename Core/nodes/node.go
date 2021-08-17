package nodes

import "github.com/hajimehoshi/ebiten/v2"

type Node interface {
	GetType() string
	GetID() uint64
	GetName() string
	Draw(screen *ebiten.Image)
	Update()
}
