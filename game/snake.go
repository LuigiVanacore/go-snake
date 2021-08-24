package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

const (
	Up = iota
	Down
	Left
	Right
	None
)

type Snake struct {
	direction int
	snakeBody []SnakePart
	size      int
	speed     int
	lives     int
	moveTime  int
	timer     int
	isDead    bool
}

func NewSnake() *Snake {
	snake := &Snake{
		snakeBody: make([]SnakePart, 1),
		moveTime:  4,
		direction: None,
	}
	snake.snakeBody[0].SetPosition(XNumInScreen/2, YNumInScreen/2)
	return snake
}

//func (s *Snake) collidesWithApple() bool {
//	return g.snakeBody[0].X == g.apple.X &&
//		g.snakeBody[0].Y == g.apple.Y
//}

func (s *Snake) collidesWithSelf() bool {
	for _, v := range s.snakeBody[1:] {
		if s.snakeBody[0].GetPosition().X == v.GetPosition().X &&
			s.snakeBody[0].GetPosition().Y == v.GetPosition().Y {
			return true
		}
	}
	return false
}

func (s *Snake) collidesWithWall() bool {
	return s.snakeBody[0].GetPosition().X < 0 ||
		s.snakeBody[0].GetPosition().Y < 0 ||
		s.snakeBody[0].GetPosition().X >= XNumInScreen ||
		s.snakeBody[0].GetPosition().Y >= YNumInScreen
}

func (s *Snake) needsToMoveSnake() bool {
	return s.timer%s.moveTime == 0
}

func (s *Snake) SetDirection(direction int) {
	s.direction = direction
}

func (s *Snake) GetDirection() int {
	return s.direction
}

func (s *Snake) reset() {
	s.moveTime = 4
	s.snakeBody = s.snakeBody[:1]
	s.snakeBody[0].SetPosition(XNumInScreen/2, YNumInScreen/2)
	s.direction = None
}

func (s *Snake) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if s.direction != Right {
			s.direction = Left
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if s.direction != Left {
			s.direction = Right
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if s.direction != Up {
			s.direction = Down
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if s.direction != Down {
			s.direction = Up
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.reset()
	}

	if s.needsToMoveSnake() {
		if s.collidesWithWall() || s.collidesWithSelf() {
			s.reset()
		}

		//if s.collidesWithApple() {
		//	g.apple.X = rand.Intn(Utils.XNumInScreen - 1)
		//	g.apple.Y = rand.Intn(Utils.YNumInScreen - 1)
		//	g.snakeBody = append(g.snakeBody, game.Position{
		//		X: g.snakeBody[len(g.snakeBody)-1].X,
		//		Y: g.snakeBody[len(g.snakeBody)-1].Y,
		//	})
		//	if len(g.snakeBody) > 10 && len(g.snakeBody) < 20 {
		//		g.level = 2
		//		g.moveTime = 3
		//	} else if len(g.snakeBody) > 20 {
		//		g.level = 3
		//		g.moveTime = 2
		//	} else {
		//		g.level = 1
		//	}
		//	g.score++
		//	if g.bestScore < g.score {
		//		g.bestScore = g.score
		//	}
		//}

		for i := int64(len(s.snakeBody)) - 1; i > 0; i-- {
			s.snakeBody[i].SetPosition(s.snakeBody[i-1].GetPosition().X,
				s.snakeBody[i-1].GetPosition().Y)
		}
		s.snakeBody[0].Move(s.direction)
	}

	s.timer++

	return nil
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, v := range s.snakeBody {
		ebitenutil.DrawRect(screen, float64(v.GetPosition().X*GridSize), float64(v.GetPosition().Y*GridSize), GridSize, GridSize, color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	}
}

//
//type SnakeLiveState struct {
//}
//
//func (s *SnakeLiveState) Execute() {
//
//}
//
//type SnakeDeadState struct {
//}
//
//func NewSnake() *Snake {
//	return &Snake{direction: None}
//}
//
//func (s *Snake) Init() {
//	s.IncreaseBody(Vector2D.Zero())
//}
//
//func (s *Snake) IncreaseBody(position Vector2D.Vector2D) {
//	snakePart := SnakePart{&ShapeNode{Image: ebiten.NewImage(blockSize, blockSize), Transform: ebiten.GeoM{}}, position}
//	snakePart.GetNode().GetImage().Fill(color.RGBA{R: 0xff, A: 0xff})
//	s.body = append(s.body, &snakePart)
//}
//
//func (s *Snake) GetBodyPart(index uint) *SnakePart {
//	return s.body[index]
//}
//
//func (s *Snake) SetDirection(direction Direction) {
//	s.direction = direction
//}
//
//func (s *Snake) GetDirection() Direction {
//	return s.direction
//}
//
//func (s *Snake) GetSpeed() int {
//	return s.speed
//}
//
//func (s *Snake) GetLives() int {
//	return s.lives
//}
//
//func (s *Snake) GetScore() int {
//	return s.score
//}
//
//func (s *Snake) IncreaseScore() {
//	s.score++
//}
//
//func (s *Snake) HasLost() bool {
//	return s.isDead
//}
//
//func (s *Snake) Reset() {
//	s.body = nil
//	transform := Vector2D.Zero()
//	s.IncreaseBody(transform)
//	s.IncreaseBody(transform)
//	s.IncreaseBody(transform)
//	s.SetDirection(None)
//	s.speed = 15
//	s.lives = 3
//	s.score = 0
//	s.isDead = false
//}
//
//func (s *Snake) GetPhysicalDirection() Direction {
//	if len(s.body) <= 1 {
//		return None
//	}
//	head := s.GetBodyPart(0)
//	neck := s.GetBodyPart(1)
//
//	if head.Position.X == neck.Position.X {
//		if head.Position.Y > neck.Position.Y {
//			return Down
//		} else {
//			return Left
//		}
//	}
//	return None
//}
//
//func (s *Snake) Draw(screen *ebiten.Image) {
//	for _, part := range s.body {
//		part.Draw(screen)
//	}
//}
//
//func (s *Snake) CheckCollision() {
//
//}
//
//func (s *Snake) Cut(segments int) {
//
//}
//
//func (s *Snake) Move() {
//	//for i := len(s.body) -1; i > 0; i-- {
//	//	s.body[i].Position = s.body[i-1].Position
//	//}
//	//for (int i = m_snakeBody.size() - 1; i > 0; --i){
//	//	m_snakeBody[i].position = m_snakeBody[i - 1].position;
//	//}
//	if s.direction == Left {
//		s.body[0].Position.X--
//	} else if s.direction == Right {
//		s.body[0].Position.X++
//	} else if s.direction == Up {
//		s.body[0].Position.Y--
//	} else if s.direction == Down {
//		s.body[0].Position.Y++
//	}
//}
//
//func (s *Snake) Update() {
//	s.Move()
//}

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
