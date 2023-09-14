package main

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) neg() Point3D {
	return Point3D{-p.X, -p.Y, -p.Z}
}

func (a Point3D) add(b Point3D) Point3D {
	return Point3D{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Point3D) sub(b Point3D) Point3D {
	return Point3D{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Point3D) taxicabDistance(b Point3D) int {
	return absInt(a.X-b.X) + absInt(a.Y-b.Y) + absInt(a.Z-b.Z)
}
