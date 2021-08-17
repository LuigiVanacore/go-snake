package Transform

type Position struct {
	X, Y float64
}

func NewPosition(x float64, y float64) *Position {
	return &Position{x, y}
}

func (p *Position) GetPosition() (float64, float64) {
	return p.X, p.Y
}

func (p *Position) SetPosition(x float64, y float64) {
	p.X = x
	p.Y = y
}

func (p *Position) Translate(x float64, y float64) {
	p.X += x
	p.Y += y
}
