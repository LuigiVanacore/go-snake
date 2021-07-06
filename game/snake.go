package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"luigi.vanacore/go-snake/Core/Vector2D"
)

type Direction int

const (
	Up = iota
	Down
	Left
	Right
	None
)

type Snake struct {
	direction Direction
	size      int
	speed     int
	lives     int
	score     int
	isDead    bool
	body      []*SnakePart
}

func NewSnake() *Snake {
	return &Snake{direction: None}
}

func (s *Snake) Init() {
	s.IncreaseBody(Vector2D.Zero())
}

func (s *Snake) IncreaseBody(position Vector2D.Vector2D) {
	snakePart := SnakePart{&ShapeNode{Image: ebiten.NewImage(blockSize, blockSize), Transform: ebiten.GeoM{}}, position}
	snakePart.GetNode().GetImage().Fill(color.RGBA{R: 0xff, A: 0xff})
	s.body = append(s.body, &snakePart)
}

func (s *Snake) GetBodyPart(index uint) *SnakePart {
	return s.body[index]
}

func (s *Snake) SetDirection(direction Direction) {
	s.direction = direction
}

func (s *Snake) GetDirection() Direction {
	return s.direction
}

func (s *Snake) GetSpeed() int {
	return s.speed
}

func (s *Snake) GetLives() int {
	return s.lives
}

func (s *Snake) GetScore() int {
	return s.score
}

func (s *Snake) IncreaseScore() {
	s.score++
}

func (s *Snake) HasLost() bool {
	return s.isDead
}

func (s *Snake) Reset() {
	s.body = nil
	transform := Vector2D.Zero()
	s.IncreaseBody(transform)
	s.IncreaseBody(transform)
	s.IncreaseBody(transform)
	s.SetDirection(None)
	s.speed = 15
	s.lives = 3
	s.score = 0
	s.isDead = false
}

func (s *Snake) GetPhysicalDirection() Direction {
	if len(s.body) <= 1 {
		return None
	}
	head := s.GetBodyPart(0)
	neck := s.GetBodyPart(1)

	if head.Position.X == neck.Position.X {
		if head.Position.Y > neck.Position.Y {
			return Down
		} else {
			return Left
		}
	}
	return None
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, part := range s.body {
		part.Draw(screen)
	}
}

func (s *Snake) CheckCollision() {

}

func (s *Snake) Cut(segments int) {

}

func (s *Snake) Move() {
	//for i := len(s.body) -1; i > 0; i-- {
	//	s.body[i].Position = s.body[i-1].Position
	//}
	//for (int i = m_snakeBody.size() - 1; i > 0; --i){
	//	m_snakeBody[i].position = m_snakeBody[i - 1].position;
	//}
	if s.direction == Left {
		s.body[0].Position.X--
	} else if s.direction == Right {
		s.body[0].Position.X++
	} else if s.direction == Up {
		s.body[0].Position.Y--
	} else if s.direction == Down {
		s.body[0].Position.Y++
	}
}

func (s *Snake) Update() {
	s.Move()
}

//Direction Snake::GetPhysicalDirection(){
//if(m_snakeBody.size() <= 1){
//return Direction::None;
//}
//
//SnakeSegment& head = m_snakeBody[0];
//SnakeSegment& neck = m_snakeBody[1];
//
//if(head.position.x == neck.position.x){
//return (head.position.y > neck.position.y
//? Direction::Down : Direction::Up);
//} else if(head.position.y == neck.position.y){
//return (head.position.x > neck.position.x
//? Direction::Right : Direction::Left);
//}
//
//return Direction::None;
//}

//Snake(int l_blockSize, Textbox* l_log);
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
//Direction GetPhysicalDirection();
//
//void Extend(); // Grow the snake.
//void Reset(); // Reset to starting position.
//
//void Move(); // Movement method.
//void Tick(); // Update method.
//void Cut(int l_segments); // Method for cutting snake.

//
//SnakeContainer m_snakeBody; // Segment vector.
//int m_size; // Size of the graphics.
//Direction m_dir; // Current direction.
//int m_speed; // Speed of the snake.
//int m_lives; // Lives.
//int m_score; // Score.
//bool m_lost; // Losing state.
//sf::RectangleShape m_bodyRect; // Shapes used in rendering.
//Textbox* m_log;
