package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Engine struct {
	window  *sdl.Window
	surface *sdl.Surface
	running bool
	input   *Input
}

type UpdateFn func(*Input, float64)
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
	e.running = true

	const secsPerUpdate = 1 / 60.0
	var current, elapsed, lag float64
	previous := float64(sdl.GetTicks()) * 0.001

	for e.running {
		current = float64(sdl.GetTicks()) * 0.001
		elapsed = current - previous
		previous = current

		if elapsed > 1.0 {
			continue
		}

		lag += elapsed

		// Handling inputs
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				e.running = false
				break
			case *sdl.KeyboardEvent:
				key := sdl.GetKeyName(t.Keysym.Sym)

				switch t.GetType() {
				case sdl.KEYDOWN:
					e.input.Keyboard.Press(key)
				case sdl.KEYUP:
					e.input.Keyboard.Release(key)
				}

			}
		}

		for lag >= secsPerUpdate {
			updateFn(e.input, secsPerUpdate)
			lag -= secsPerUpdate
		}

		e.surface.FillRect(nil, 0)
		drawFn(e.surface)
		e.window.UpdateSurface()
	}
}

func NewEngine() *Engine {
	return &Engine{
		window:  nil,
		surface: nil,
		running: false,
		input:   NewInput(),
	}
}
