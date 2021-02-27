package main

import (
	"github.com/gabrielEscame/go-engine/engine"
	"github.com/gabrielEscame/go-engine/pong"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	ball := pong.NewBall(50, 50, 10)
	player := pong.NewPlayer()
	e := engine.NewEngine()

	e.Setup()
	e.Loop(func(i *engine.Input) {
		ball.Update(i)
		player.Update(i)
	}, func(s *sdl.Surface) {
		ball.Draw(s)
		player.Draw(s)
	})

}
