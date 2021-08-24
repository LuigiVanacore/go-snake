package game

//
//const blockSize = 16
//
//type World struct {
//	windowsWidth int
//	windowHeigth int
//	snake        *Snake
//	apple        *ShapeNode
//}
//
//func NewWorld(screenWidth int, screenHeight int) *World {
//	snake := New()
//	return &World{windowsWidth: screenWidth, windowHeigth: screenHeight, snake: snake}
//}
//
//func (w *World) Init() {
//	w.snake.Init()
//}
//
//func (w *World) Update() {
//	w.snake.Update()
//	//if(l_player.GetPosition() == m_item){
//	//	l_player.Extend();
//	//	l_player.IncreaseScore();
//	//	RespawnApple();
//	//}
//	//
//	//int gridSize_x = m_windowSize.x / m_blockSize;
//	//int gridSize_y = m_windowSize.y / m_blockSize;
//	//
//	//if(l_player.GetPosition().x <= 0 ||
//	//	l_player.GetPosition().y <= 0 ||
//	//	l_player.GetPosition().x >= gridSize_x - 1 ||
//	//	l_player.GetPosition().y >= gridSize_y - 1)
//	//{
//	//	l_player.Lose();
//	//}
//}
//
//func (w *World) Draw(screen *ebiten.Image) {
//	w.snake.Draw(screen)
//}
//
//func (w *World) RespawnApple() {
//	maxX := (w.windowsWidth / blockSize) - 2
//	maxY := (w.windowHeigth / blockSize) - 2
//	positionX := rand.Intn(maxX-1) + 1
//	positionY := rand.Intn(maxY-1) + 1
//	w.apple.GetTransform().Apply(float64(positionX*blockSize), float64(positionY*blockSize))
//}
//
////
////class World{
////public:
////World(sf::Vector2u l_windSize);
////~World();
////
////int GetBlockSize();
////
////void RespawnApple();
////
////void Update(Snake& l_player);
////void Render(sf::RenderWindow& l_window);
////private:
////sf::Vector2u m_windowSize;
////sf::Vector2i m_item;
////int m_blockSize;
////
////sf::CircleShape m_appleShape;
////sf::RectangleShape m_bounds[4];
////};
