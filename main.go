package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var node *ShapeNode

const (
	Up = iota
	Down
	Left
	Right
)

type ShapeNode struct {
	Image     *ebiten.Image
	Transform ebiten.GeoM
}

func (s *ShapeNode) getImage() *ebiten.Image {
	return s.Image
}

type SnakePart struct {
	BodyRect *ShapeNode
}

func (s *SnakePart) getNode() *ShapeNode {
	return s.BodyRect
}

type Snake struct {
	direction int
	size      int
	speed     int
	lives     int
	scores    int
	isDead    bool
	body      []*SnakePart
}

func (s *Snake) Init() {
	s.IncreaseBody()
}

func (s *Snake) IncreaseBody() {
	snakePart := SnakePart{&ShapeNode{Image: ebiten.NewImage(50, 50), Transform: ebiten.GeoM{}}}
	snakePart.getNode().getImage().Fill(color.RGBA{0xff, 0, 0, 0xff})
	s.body = append(s.body, &snakePart)
}

func (s *Snake) getBodyPart() *SnakePart {
	return s.body[0]
}

//var image *ebiten.Image

type Game struct {
	snake Snake
}

//
//func (g *Game) Update() error {
//return nil
//}
//
//func (g *Game) Draw(screen *ebiten.Image) {
//	op := &ebiten.DrawImageOptions{}
//	screen.DrawImage(image, op)
//
//}
//
//
//func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
//	return screenWidth, screenHeight
//}

//class Snake{
//public:
//Snake(int l_blockSize);
//~Snake();
//
//// Helper methods.
//void SetDirection(Direction l_dir);
//Direction GetDirection();
//int GetSpeed();
//sf::Vector2i GetPosition();
//int GetLives();
//int GetScore();
//void IncreaseScore();
//bool HasLost();
//
//void Lose(); // Handle losing here.
//void ToggleLost();
//
//void Extend(); // Grow the snake.
//void Reset(); // Reset to starting position.
//
//void Move(); // Movement method.
//void Tick(); // Update method.
//void Cut(int l_segments); // Method for cutting snake.
//void Render(sf::RenderWindow& l_window);
//private:
//void CheckCollision(); // Checking for collisions.
//
//SnakeContainer m_snakeBody; // Segment vector.
//int m_size; // Size of the graphics.

//};

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

func main() {
	game := &Game{}
	game.Init()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("go-snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
