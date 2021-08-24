package Collision

type CollisionShape interface {
	IsColliding(collision CollisionShape) (error, bool)
	WillCollide(collision CollisionShape, dx float64, dy float64) bool
}
