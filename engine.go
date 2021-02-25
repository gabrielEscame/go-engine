package main

import "github.com/veandco/go-sdl2/sdl"

type Engine struct {
	window  *sdl.Window
	surface *sdl.Surface
	running bool
}

type UpdateFn func()
type DrawFn func(*sdl.Surface)

func (e *Engine) Setup() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(
		"Engine",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		256, 240,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	e.window = window
	e.surface = surface
}

func (e *Engine) Loop(updateFn UpdateFn, drawFn DrawFn) {
	for e.running {

		// Handling inputs
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				e.running = false
				break
			}
		}

		updateFn()

		// e.surface.FillRect(nil, 0)
		drawFn(e.surface)
		e.window.UpdateSurface()

	}
}

func NewEngine() *Engine {
	return &Engine{
		window:  nil,
		surface: nil,
		running: true,
	}
}
