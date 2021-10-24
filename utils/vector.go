package utils

type Vector struct {
	X, Y int
}

func (v Vector) Add(av Vector) Vector {
	return Vector{
		X: v.X + av.X,
		Y: v.Y + av.Y,
	}
}
