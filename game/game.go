package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world *World
}

func NewGame(screenWeight int, screenHeight int) *Game {
	world := NewWorld(screenWeight, screenHeight)
	return &Game{world: world}
}

func (g *Game) Init() {
	g.world.Init()
}

func (g *Game) Update() error {
	g.HandleInput()
	g.world.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) HandleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.world.snake.SetDirection(Up)
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.world.snake.SetDirection(Down)
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.world.snake.SetDirection(Right)
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.world.snake.SetDirection(Left)
	}

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
}
