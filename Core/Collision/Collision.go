package Collision

type Collision interface {
	IsColliding(collision Collision) bool
	WillCollide(collision Collision, dx float64, dy float64) bool
}
