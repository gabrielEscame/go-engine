package main

import "github.com/veandco/go-sdl2/sdl"

func main() {

	ball := NewBall(50, 50, 10)

	engine := NewEngine()
	engine.Setup()
	engine.Loop(func() {
		ball.Update()
	}, func(s *sdl.Surface) {
		ball.Draw(s)
	})

}
