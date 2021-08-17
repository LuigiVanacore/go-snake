package nodes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"luigi.vanacore/go-snake/Core/Shapes"
)

type NodeDrawRect struct {
	id       uint64
	name     string
	nodeType string
	rect     Shapes.Rect
	color    color.RGBA
}

func NewNodeDrawRect() *NodeDrawRect {
	return &NodeDrawRect{nodeType: "DrawRect"}
}

func (n *NodeDrawRect) GetType() string {
	return n.nodeType
}

func (n *NodeDrawRect) GetID() uint64 {
	return n.id
}

func (n *NodeDrawRect) GetName() string {
	return n.name
}

func (n *NodeDrawRect) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, n.rect.X, n.rect.Y, n.rect.Width, n.rect.Height, n.color)
}

func (n *NodeDrawRect) Update() {

}
