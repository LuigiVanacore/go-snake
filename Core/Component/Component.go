package Component

type Component interface {
	GetType()
	GetID()
	Update()
	Draw()
}
