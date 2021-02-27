package engine

import "github.com/veandco/go-sdl2/sdl"

type Entity interface {
	Update(*Input)
	Draw(g *sdl.Surface)
}
