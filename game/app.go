package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	titleWindow  = "Snake"
)

type App struct {
}

func init() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle(titleWindow)
}

func (app *App) Run() {
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
