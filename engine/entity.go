package engine

import (
	"github.com/gabrielEscame/go-engine/physics"
	"github.com/veandco/go-sdl2/sdl"
)

type Entity interface {
	Update(*Input, float64)
	Draw(*sdl.Surface)
}

type CollidableEntity interface {
	GetShape() *physics.SquareShape
	OnCollisionEnter(*CollisionInfo)
}
