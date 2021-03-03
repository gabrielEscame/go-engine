package main

import (
	"fmt"
	"testing"

	"github.com/gabrielEscame/go-engine/engine"
	"github.com/gabrielEscame/go-engine/pong"
)

func TestTest(t *testing.T) {
	fmt.Println("opa")

	list := []engine.CollidableEntity{
		pong.NewBall(0, 0, 10),
		pong.NewBall(0, 0, 10),
		pong.NewBall(10, 110, 10),
		pong.NewPlayer(),
	}

	engine.CollisionChecker(list)
}
