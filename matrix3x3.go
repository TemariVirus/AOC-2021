package main

import "strconv"

type Matrix3x3 struct {
	data [3][3]int
}

func (left Matrix3x3) multiply(right Matrix3x3) Matrix3x3 {
	data := [3][3]int{}
	for i := range data {
		for j := range data[i] {
			for k := range data[i] {
				data[i][j] += left.data[i][k] * right.data[k][j]
			}
		}
	}
	return Matrix3x3{data}
}

func (mat Matrix3x3) transform(p Point3D) Point3D {
	return Point3D{
		p.X*mat.data[0][0] + p.Y*mat.data[0][1] + p.Z*mat.data[0][2],
		p.X*mat.data[1][0] + p.Y*mat.data[1][1] + p.Z*mat.data[1][2],
		p.X*mat.data[2][0] + p.Y*mat.data[2][1] + p.Z*mat.data[2][2],
	}
}

func (m Matrix3x3) toString() string {
	return "[[" + join(m.data[:], "],\n[", func(row [3]int) string {
		return join(row[:], ", ", strconv.Itoa)
	}) + "]]"
}
