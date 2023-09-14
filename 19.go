package main

import (
	"strconv"
	"strings"
)

type Scanner struct {
	Position  *Point3D
	Direction *Matrix3x3
	Beacons   []Point3D
}

var identity = Matrix3x3{[3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}}
var x = Matrix3x3{[3][3]int{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}}}
var y = Matrix3x3{[3][3]int{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}}}
var z = Matrix3x3{[3][3]int{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}}}

var directions = [24]Matrix3x3{
	identity,
	identity.multiply(x),
	identity.multiply(x).multiply(x),
	identity.multiply(x).multiply(x).multiply(x),
	identity.multiply(y),
	identity.multiply(y).multiply(x),
	identity.multiply(y).multiply(x).multiply(x),
	identity.multiply(y).multiply(x).multiply(x).multiply(x),
	identity.multiply(y).multiply(y),
	identity.multiply(y).multiply(y).multiply(x),
	identity.multiply(y).multiply(y).multiply(x).multiply(x),
	identity.multiply(y).multiply(y).multiply(x).multiply(x).multiply(x),
	identity.multiply(y).multiply(y).multiply(y),
	identity.multiply(y).multiply(y).multiply(y).multiply(x),
	identity.multiply(y).multiply(y).multiply(y).multiply(x).multiply(x),
	identity.multiply(y).multiply(y).multiply(y).multiply(x).multiply(x).multiply(x),
	identity.multiply(z),
	identity.multiply(z).multiply(z).multiply(z),
	identity.multiply(y).multiply(z),
	identity.multiply(y).multiply(z).multiply(z).multiply(z),
	identity.multiply(y).multiply(y).multiply(z),
	identity.multiply(y).multiply(y).multiply(z).multiply(z).multiply(z),
	identity.multiply(y).multiply(y).multiply(y).multiply(z),
	identity.multiply(y).multiply(y).multiply(y).multiply(z).multiply(z).multiply(z),
}

func parseInput19(input string) []Scanner {
	scans := []Scanner{}
	for _, part := range strings.Split(input, "\n\n") {
		lines := strings.Split(part, "\n")[1:]
		beacons := apply(lines, func(line string) Point3D {
			coords := apply(strings.Split(line, ","), func(s string) int {
				return unwrap(strconv.Atoi(s))
			})
			return Point3D{coords[0], coords[1], coords[2]}
		})
		scans = append(scans, Scanner{nil, nil, beacons})
	}

	scans[0].Position = &Point3D{0, 0, 0}
	scans[0].Direction = &identity
	return scans
}

func beaconsAbsoluteCoords(ref Scanner, beacons []Point3D) []Point3D {
	for _, p1 := range ref.Beacons {
		scans1 := makeSetFrom(apply(ref.Beacons, p1.sub))
		for _, p2 := range beacons[:len(beacons)-11] {
			scans2 := makeSetFrom(apply(beacons, p2.sub))
			if scans1.intersect(scans2).len() >= 12 {
				scans := apply(beacons, p2.neg().add)
				scans = apply(scans, p1.add)
				return apply(scans, ref.Position.add)
			}
		}
	}
	return nil
}

func findScannersAndBeacons(scanners []Scanner) ([]Scanner, Set[Point3D]) {
	searched := [][]bool{}
	for i := 0; i < len(scanners); i++ {
		searched = append(searched, make([]bool, len(scanners)))
	}

	beacons := makeSetFrom(scanners[0].Beacons)
	for anyTrue(scanners, func(s Scanner) bool { return s.Position == nil }) {
		for i, ref := range scanners {
			if ref.Position == nil {
				continue
			}

			for j, scanner := range scanners {
				if searched[i][j] || scanner.Position != nil {
					continue
				}

				searched[i][j] = true
				for _, d := range directions {
					scans := apply(scanner.Beacons, d.transform)
					abs_scans := beaconsAbsoluteCoords(ref, scans)
					if abs_scans == nil {
						continue
					}
					beacons = beacons.union(makeSetFrom(abs_scans))
					pos := abs_scans[0].sub(scans[0])
					scanners[j].Position = &pos
					scanners[j].Direction = &d
					scanners[j].Beacons = scans
					break
				}
			}
		}
	}

	return scanners, beacons
}

func solution19Part1(input string) int {
	scanners := parseInput19(input)
	_, beacons := findScannersAndBeacons(scanners)
	return beacons.len()
}

func solution19Part2(input string) int {
	scanners := parseInput19(input)
	scanners, _ = findScannersAndBeacons(scanners)

	positions := apply(scanners, func(s Scanner) Point3D { return *s.Position })
	max_dist := 0
	for i, pos1 := range positions {
		for _, pos2 := range positions[i+1:] {
			max_dist = max(max_dist, pos1.taxicabDistance(pos2))
		}
	}

	return max_dist
}
