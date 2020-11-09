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

type SnakePart struct {

}


type Snake struct{
	size int
	speed int
	lives int
	scores int
	isDead bool
	BodyRect *ebiten.Image
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
//Direction m_dir; // Current direction.
//int m_speed; // Speed of the snake.
//int m_lives; // Lives.
//int m_score; // Score.
//bool m_lost; // Losing state.
//sf::RectangleShape m_bodyRect; // Shape used in rendering.
//};





func (g *Game) Update() error {
	return nil
}



func (g *Game) Init() {
	g.snake.BodyRect = ebiten.NewImage(50,50)
	g.snake.BodyRect.Fill(color.RGBA{0xff, 0, 0, 0xff})
	//g.snake.BodyRect.Fill(color.RGBA{0xff, 0, 0, 0xff} )
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.snake.BodyRect, nil)

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
	}}