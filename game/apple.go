package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"luigi.vanacore/go-snake/Core/Transform"
)

type Apple struct {
	Position Transform.Position
}

func NewApple(x, y float64) *Apple {
	return &Apple{Position: Transform.Position{X: x, Y: y}}
}

func (a *Apple) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, a.Position.X*GridSize, a.Position.Y*GridSize, GridSize, GridSize, color.RGBA{R: 0xFF, A: 0xff})
}
