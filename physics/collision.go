package physics

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

func AABB(a, b SquareShape) bool {
	return (a.X < b.X+b.W &&
		a.X+a.W > b.X &&
		a.Y < b.Y+b.H &&
		a.Y+a.H > b.Y)
}
