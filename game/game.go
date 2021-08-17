package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"math/rand"
	"time"
)

const (
	dirNone = iota
	dirLeft
	dirRight
	dirDown
	dirUp
)

const (
	gridSize     = 10
	xNumInScreen = ScreenWidth / gridSize
	yNumInScreen = ScreenHeight / gridSize
)

type Position struct {
	X int
	Y int
}

type Game struct {
	moveDirection int
	snakeBody     []Position
	apple         Position
	timer         int
	moveTime      int
	score         int
	bestScore     int
	level         int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame() *Game {
	g := &Game{
		apple:     Position{X: 3 * gridSize, Y: 3 * gridSize},
		moveTime:  4,
		snakeBody: make([]Position, 1),
	}
	g.snakeBody[0].X = xNumInScreen / 2
	g.snakeBody[0].Y = yNumInScreen / 2
	return g
}

func (g *Game) collidesWithApple() bool {
	return g.snakeBody[0].X == g.apple.X &&
		g.snakeBody[0].Y == g.apple.Y
}

func (g *Game) collidesWithSelf() bool {
	for _, v := range g.snakeBody[1:] {
		if g.snakeBody[0].X == v.X &&
			g.snakeBody[0].Y == v.Y {
			return true
		}
	}
	return false
}

func (g *Game) collidesWithWall() bool {
	return g.snakeBody[0].X < 0 ||
		g.snakeBody[0].Y < 0 ||
		g.snakeBody[0].X >= xNumInScreen ||
		g.snakeBody[0].Y >= yNumInScreen
}

func (g *Game) needsToMoveSnake() bool {
	return g.timer%g.moveTime == 0
}

func (g *Game) reset() {
	g.apple.X = 3 * gridSize
	g.apple.Y = 3 * gridSize
	g.moveTime = 4
	g.snakeBody = g.snakeBody[:1]
	g.snakeBody[0].X = xNumInScreen / 2
	g.snakeBody[0].Y = yNumInScreen / 2
	g.score = 0
	g.level = 1
	g.moveDirection = dirNone
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.moveDirection != dirRight {
			g.moveDirection = dirLeft
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.moveDirection != dirLeft {
			g.moveDirection = dirRight
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.moveDirection != dirUp {
			g.moveDirection = dirDown
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.moveDirection != dirDown {
			g.moveDirection = dirUp
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.reset()
	}

	if g.needsToMoveSnake() {
		if g.collidesWithWall() || g.collidesWithSelf() {
			g.reset()
		}

		if g.collidesWithApple() {
			g.apple.X = rand.Intn(xNumInScreen - 1)
			g.apple.Y = rand.Intn(yNumInScreen - 1)
			g.snakeBody = append(g.snakeBody, Position{
				X: g.snakeBody[len(g.snakeBody)-1].X,
				Y: g.snakeBody[len(g.snakeBody)-1].Y,
			})
			if len(g.snakeBody) > 10 && len(g.snakeBody) < 20 {
				g.level = 2
				g.moveTime = 3
			} else if len(g.snakeBody) > 20 {
				g.level = 3
				g.moveTime = 2
			} else {
				g.level = 1
			}
			g.score++
			if g.bestScore < g.score {
				g.bestScore = g.score
			}
		}

		for i := int64(len(g.snakeBody)) - 1; i > 0; i-- {
			g.snakeBody[i].X = g.snakeBody[i-1].X
			g.snakeBody[i].Y = g.snakeBody[i-1].Y
		}
		switch g.moveDirection {
		case dirLeft:
			g.snakeBody[0].X--
		case dirRight:
			g.snakeBody[0].X++
		case dirDown:
			g.snakeBody[0].Y++
		case dirUp:
			g.snakeBody[0].Y--
		}
	}

	g.timer++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.snakeBody {
		ebitenutil.DrawRect(screen, float64(v.X*gridSize), float64(v.Y*gridSize), gridSize, gridSize, color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	}
	ebitenutil.DrawRect(screen, float64(g.apple.X*gridSize), float64(g.apple.Y*gridSize), gridSize, gridSize, color.RGBA{0xFF, 0x00, 0x00, 0xff})

	if g.moveDirection == dirNone {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Press up/down/left/right to start"))
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f Level: %d Score: %d Best Score: %d", ebiten.CurrentFPS(), g.level, g.score, g.bestScore))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

//type Game struct {
//	world *World
//}
//
//func NewGame(screenWeight int, screenHeight int) *Game {
//	world := NewWorld(screenWeight, screenHeight)
//	return &Game{world: world}
//}
//
//func (g *Game) Init() {
//	g.world.Init()
//}
//
//func (g *Game) Update() error {
//	g.HandleInput()
//	g.world.Update()
//	return nil
//}
//
//func (g *Game) Draw(screen *ebiten.Image) {
//	g.world.Draw(screen)
//
//}
//
//func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	return 320, 240
//}
//
//func (g *Game) HandleInput() {
//	if ebiten.IsKeyPressed(ebiten.KeyUp) {
//		g.world.snake.SetDirection(Up)
//	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
//		g.world.snake.SetDirection(Down)
//	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
//		g.world.snake.SetDirection(Right)
//	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
//		g.world.snake.SetDirection(Left)
//	}

//if(sf::Keyboard::isKeyPressed(sf::Keyboard::Up)
//&& m_snake.GetPhysicalDirection() != Direction::Down){
//	m_snake.SetDirection(Direction::Up);
//} else if(sf::Keyboard::isKeyPressed(sf::Keyboard::Down)
//&& m_snake.GetPhysicalDirection() != Direction::Up){
//	m_snake.SetDirection(Direction::Down);
//} else if(sf::Keyboard::isKeyPressed(sf::Keyboard::Left)
//&& m_snake.GetPhysicalDirection() != Direction::Right){
//	m_snake.SetDirection(Direction::Left);
//} else if(sf::Keyboard::isKeyPressed(sf::Keyboard::Right)
//&& m_snake.GetPhysicalDirection() != Direction::Left){
//	m_snake.SetDirection(Direction::Right);
//}
//}
