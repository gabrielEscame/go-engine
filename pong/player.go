package pong

import (
	"github.com/gabrielEscame/go-engine/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	x float64
	y float64
	w float64
	h float64
}

func (p *Player) Update(i *engine.Input, dt float64) {
	if i.Keyboard.IsPressed("W") && p.y > 0 {
		p.y -= 100 * dt
	}
	if i.Keyboard.IsPressed("S") && p.y+p.h < 240 {
		p.y += 100 * dt
	}
}

func (p *Player) Draw(g *sdl.Surface) {
	rect := sdl.Rect{
		X: int32(p.x),
		Y: int32(p.y),
		W: int32(p.w),
		H: int32(p.h),
	}
	g.FillRect(&rect, 0xffffffff)
}

func NewPlayer() *Player {
	return &Player{
		x: 10,
		y: 110,
		w: 10,
		h: 50,
	}
}
