package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
	"time"
)

const (
	GridSize     = 10
	XNumInScreen = ScreenWidth / GridSize
	YNumInScreen = ScreenHeight / GridSize
)

type Position struct {
	X int
	Y int
}

type Game struct {
	snake     *Snake
	apple     Position
	score     int
	bestScore int
	level     int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame() *Game {
	g := &Game{
		apple: Position{X: 3 * GridSize, Y: 3 * GridSize},
	}
	g.snake = NewSnake()
	return g
}

//
//func (g *Game) collidesWithApple() bool {
//	return g.snakeBody[0].X == g.apple.X &&
//		g.snakeBody[0].Y == g.apple.Y
//}
//
//func (g *Game) collidesWithSelf() bool {
//	for _, v := range g.snakeBody[1:] {
//		if g.snakeBody[0].X == v.X &&
//			g.snakeBody[0].Y == v.Y {
//			return true
//		}
//	}
//	return false
//}
//
//func (g *Game) collidesWithWall() bool {
//	return g.snakeBody[0].X < 0 ||
//		g.snakeBody[0].Y < 0 ||
//		g.snakeBody[0].X >= xNumInScreen ||
//		g.snakeBody[0].Y >= yNumInScreen
//}
//
//func (g *Game) needsToMoveSnake() bool {
//	return g.timer%g.moveTime == 0
//}
//
//func (g *Game) SetDirection(direction int) {
//	g.moveDirection = direction
//}
//
//func (g *Game) GetDirection() int {
//	return g.moveDirection
//}
//
//func (g *Game) reset() {
//	g.apple.X = 3 * Utils.GridSize
//	g.apple.Y = 3 * Utils.GridSize
//	g.moveTime = 4
//	g.snakeBody = g.snakeBody[:1]
//	g.snakeBody[0].X = xNumInScreen / 2
//	g.snakeBody[0].Y = yNumInScreen / 2
//	g.score = 0
//	g.level = 1
//	g.moveDirection = InputManager.dirNone
//}

func (g *Game) Update() error {
	g.snake.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.snake.Draw(screen)
	ebitenutil.DrawRect(screen, float64(g.apple.X*GridSize), float64(g.apple.Y*GridSize), GridSize, GridSize, color.RGBA{0xFF, 0x00, 0x00, 0xff})

	if g.snake.GetDirection() == None {
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
