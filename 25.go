package main

import (
	"strings"
)

type SeaCucumber int8

const (
	None SeaCucumber = iota
	East
	South
)

func parseInput25(input string) [][]SeaCucumber {
	floor := [][]SeaCucumber{}
	for _, line := range strings.Split(input, "\n") {
		row := []SeaCucumber{}
		for _, char := range line {
			switch char {
			case '.':
				row = append(row, None)
			case '>':
				row = append(row, East)
			case 'v':
				row = append(row, South)
			}
		}
		floor = append(floor, row)
	}
	return floor
}

func moveCucumbers(floor [][]SeaCucumber) (new_floor [][]SeaCucumber, moved int) {
	new_floor = make([][]SeaCucumber, 0, len(floor))
	for _, row := range floor {
		new_row := make([]SeaCucumber, 0, len(row))
		for x, cucumber := range row {
			switch cucumber {
			case None:
				x = (x + len(row) - 1) % len(row)
				if row[x] == East {
					new_row = append(new_row, East)
				} else {
					new_row = append(new_row, None)
				}
			case East:
				x = (x + 1) % len(row)
				if row[x] == None {
					new_row = append(new_row, None)
					moved++
				} else {
					new_row = append(new_row, East)
				}
			case South:
				new_row = append(new_row, South)
			}
		}
		new_floor = append(new_floor, new_row)
	}
	floor = new_floor

	new_floor = make([][]SeaCucumber, 0, len(floor))
	for y, row := range floor {
		new_row := make([]SeaCucumber, 0, len(row))
		for x, cucumber := range row {
			switch cucumber {
			case None:
				y := (y + len(floor) - 1) % len(floor)
				if floor[y][x] == South {
					new_row = append(new_row, South)
				} else {
					new_row = append(new_row, None)
				}
			case East:
				new_row = append(new_row, East)
			case South:
				y := (y + 1) % len(floor)
				if floor[y][x] == None {
					moved++
					new_row = append(new_row, None)
				} else {
					new_row = append(new_row, South)
				}
			}
		}
		new_floor = append(new_floor, new_row)
	}

	return new_floor, moved
}

func solution25Part1(input string) int {
	floor := parseInput25(input)
	moved := 0
	for i := 1; ; i++ {
		floor, moved = moveCucumbers(floor)
		if moved == 0 {
			return i
		}
	}
}
