package main

type Point struct {
	X, Y int
}

func taxicabDistance(a, b Point) int {
	return absInt(a.X-b.X) + absInt(a.Y-b.Y)
}
