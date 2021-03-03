package pong

import (
	"fmt"
	"math/rand"

	"github.com/gabrielEscame/go-engine/engine"
	"github.com/gabrielEscame/go-engine/physics"
	"github.com/veandco/go-sdl2/sdl"
)

type Ball struct {
	x      float64
	y      float64
	radius float64
	dirX   float64
	dirY   float64
}

func (b *Ball) Update(i *engine.Input, dt float64) {
	b.x += b.dirX * dt * 100
	b.y += b.dirY * dt * 100

	if b.x+b.radius*2 > 256 || b.x < 0 {
		b.dirX *= -1
	}

	if b.y+b.radius*2 > 240 || b.y < 0 {
		b.dirY *= -1
	}
}

func (b *Ball) Draw(g *sdl.Surface) {
	rect := sdl.Rect{
		X: int32(b.x),
		Y: int32(b.y),
		W: int32(b.radius * 2),
		H: int32(b.radius * 2),
	}
	g.FillRect(&rect, uint32(rand.Intn(0xffffffff)))
}

func (b *Ball) GetShape() *physics.SquareShape {
	return &physics.SquareShape{
		X: b.x,
		Y: b.y,
		W: 10,
		H: 10,
	}
	// return &physics.CircleShape{
	// 	X:      b.x,
	// 	Y:      b.y,
	// 	Radius: b.radius,
	// }
}

func (b *Ball) OnCollisionEnter(collisionInfo *engine.CollisionInfo) {
	fmt.Println("Bola colidiu com alguma coisa!")
	b.dirX *= -1
}

func NewBall(x, y, radius float64) *Ball {
	return &Ball{
		x:      x,
		y:      y,
		radius: radius,
		dirX:   1,
		dirY:   1,
	}
}
