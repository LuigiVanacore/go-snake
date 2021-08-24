package Transform

type Transform interface {
	GetPosition() (float64, float64)
	SetPosition(float64, float64)
	Translate(float64, float64)
}
