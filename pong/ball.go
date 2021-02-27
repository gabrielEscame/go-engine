package pong

import (
	"math/rand"

	"github.com/gabrielEscame/go-engine/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type Ball struct {
	x      float64
	y      float64
	radius float64
	dirX   float64
	dirY   float64
}

func (b *Ball) Update(i *engine.Input) {
	b.x += b.dirX
	b.y += b.dirY

	if b.x > 256 || b.x < 0 {
		b.dirX *= -1
	}

	if b.y > 240 || b.y < 0 {
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

func NewBall(x, y, radius float64) *Ball {
	return &Ball{
		x:      x,
		y:      y,
		radius: radius,
		dirX:   1,
		dirY:   1,
	}
}
