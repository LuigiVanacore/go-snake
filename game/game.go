package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	snake Snake
}

func (g *Game) Init() {
	g.snake.Init()
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.snake.getBodyPart().getNode().getImage(), nil)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
