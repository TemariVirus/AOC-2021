package main

import (
	"strings"
)

func solution11Part1(input string) int {
	energies := parseEnergies(input)
	count := 0
	for i := 0; i < 100; i++ {
		for y, row := range energies {
			for x := range row {
				energies[y][x]++
			}
		}

		flashed := makeSet[Point](0)
		for y, row := range energies {
			for x, e := range row {
				if e > 9 {
					if !flashed.contains(Point{x, y}) {
						count += octoFlash(energies, Point{x, y}, flashed)
					}
				}
			}
		}
	}

	return count
}

func parseEnergies(input string) [][]int {
	heights := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		heights = append(heights, apply([]rune(line), func(char rune) int {
			return int(char) - int('0')
		}))
	}
	return heights
}

func octoFlash(energies [][]int, pos Point, flashed Set[Point]) int {
	flashed.add(pos)

	count := 1
	for dy := -1; dy <= 1; dy++ {
		y := pos.Y + dy
		if y < 0 || y >= len(energies) {
			continue
		}

		for dx := -1; dx <= 1; dx++ {
			x := pos.X + dx
			if x < 0 || x >= len(energies[y]) {
				continue
			}
			if flashed.contains(Point{x, y}) {
				continue
			}

			energies[y][x]++
			if energies[y][x] > 9 {
				count += octoFlash(energies, Point{x, y}, flashed)
			}
		}
	}

	energies[pos.Y][pos.X] = 0
	return count
}

func solution11Part2(input string) int {
	energies := parseEnergies(input)
	i := 0
	count := 0
	for ; count < len(energies)*len(energies[0]); i++ {
		count = 0

		for y, row := range energies {
			for x := range row {
				energies[y][x]++
			}
		}

		flashed := makeSet[Point](0)
		for y, row := range energies {
			for x, e := range row {
				if e > 9 {
					if !flashed.contains(Point{x, y}) {
						count += octoFlash(energies, Point{x, y}, flashed)
					}
				}
			}
		}
	}

	return i
}
