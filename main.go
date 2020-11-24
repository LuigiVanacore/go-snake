package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"luigi.vanacore/go-snake/game"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

//
//func (g *game) Update() error {
//return nil
//}
//
//func (g *game) Draw(screen *ebiten.Image) {
//	op := &ebiten.DrawImageOptions{}
//	screen.DrawImage(image, op)
//
//}
//
//
//func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
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

func main() {
	game := game.NewGame(screenWidth, screenHeight)
	game.Init()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("go-snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
