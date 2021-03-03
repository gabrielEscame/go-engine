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

	entities := []engine.CollidableEntity{
		ball,
		player,
	}

	e.Setup()
	e.Loop(func(i *engine.Input, dt float64) {
		engine.CollisionChecker(entities)
		ball.Update(i, dt)
		player.Update(i, dt)
	}, func(s *sdl.Surface) {
		ball.Draw(s)
		player.Draw(s)
	})

}
