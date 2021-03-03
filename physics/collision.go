package physics

import "math"

type SquareShape struct {
	X float64
	Y float64
	W float64
	H float64
}

type CircleShape struct {
	X      float64
	Y      float64
	Radius float64
}

func DistanceTo(x1, y1, x2, y2 float64) float64 {
	deltaX := x2 - x1
	deltaY := y2 - y1
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func AABB(a, b *SquareShape) bool {
	return (a.X < b.X+b.W &&
		a.X+a.W > b.X &&
		a.Y < b.Y+b.H &&
		a.Y+a.H > b.Y)
}

func CircleCollision(a, b *CircleShape) bool {
	totalRadiusSquared := a.Radius + b.Radius
	totalRadiusSquared = totalRadiusSquared * totalRadiusSquared

	return DistanceTo(a.X, a.Y, b.X, b.Y) < totalRadiusSquared
}

func CircleSquareCollision(a *CircleShape, b *SquareShape) bool {
	var cX, cY float64

	if a.X < b.X {
		cX = b.X
	} else if a.X > b.X+b.W {
		cX = b.X + b.W
	} else {
		cX = a.X
	}

	if a.Y < b.Y {
		cY = b.Y
	} else if a.Y > b.Y+b.H {
		cY = b.Y + b.H
	} else {
		cY = a.Y
	}

	return DistanceTo(a.X, a.Y, cX, cY) < a.Radius*a.Radius
}
