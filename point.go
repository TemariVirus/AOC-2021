package main

type Point struct {
	X, Y int
}

func taxicab_distance(a, b Point) int {
	return abs_int(a.X-b.X) + abs_int(a.Y-b.Y)
}
