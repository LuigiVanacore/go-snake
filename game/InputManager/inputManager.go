package InputManager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	game "luigi.vanacore/go-snake/game"
)

const (
	dirNone = iota
	dirLeft
	dirRight
	dirDown
	dirUp
)

type InputManager struct {
	game *game.Game
}

func New(game *game.Game) *InputManager {
	return &InputManager{
		game: game,
	}
}

func (i *InputManager) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if i.game.GetDirection() != dirRight {
			i.game.SetDirection(dirLeft)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if i.game.GetDirection() != dirLeft {
			i.game.SetDirection(dirRight)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if i.game.GetDirection() != dirUp {
			i.game.SetDirection(dirDown)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if i.game.GetDirection() != dirDown {
			i.game.SetDirection(dirUp)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		i.game.Reset()
	}
}
