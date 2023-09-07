package main

import (
	"strconv"
	"strings"
)

func solution_5_1(input string) int {
	lines := make([][4]int, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		p1 := strings.Split(parts[0], ",")
		p2 := strings.Split(parts[1], ",")
		lines = append(lines, [4]int{
			unwrap(strconv.Atoi(p1[0])),
			unwrap(strconv.Atoi(p1[1])),
			unwrap(strconv.Atoi(p2[0])),
			unwrap(strconv.Atoi(p2[1])),
		})
	}

	min_x, max_x, min_y, max_y := 0, 0, 0, 0
	for _, line := range lines {
		min_x = min(min_x, line[0], line[2])
		max_x = max(max_x, line[0], line[2])
		min_y = min(min_y, line[1], line[3])
		max_y = max(max_y, line[1], line[3])
	}

	width, height := max_x-min_x+1, max_y-min_y+1
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	for _, line := range lines {
		if line[0] == line[2] {
			for y := min(line[1], line[3]); y <= max(line[1], line[3]); y++ {
				grid[y-min_y][line[0]-min_x]++
			}
		}
		if line[1] == line[3] {
			for x := min(line[0], line[2]); x <= max(line[0], line[2]); x++ {
				grid[line[1]-min_y][x-min_x]++
			}
		}
	}

	count := 0
	for _, line := range grid {
		for _, val := range line {
			if val >= 2 {
				count++
			}
		}
	}
	return count
}

func solution_5_2(input string) int {
	lines := make([][4]int, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		p1 := strings.Split(parts[0], ",")
		p2 := strings.Split(parts[1], ",")
		lines = append(lines, [4]int{
			unwrap(strconv.Atoi(p1[0])),
			unwrap(strconv.Atoi(p1[1])),
			unwrap(strconv.Atoi(p2[0])),
			unwrap(strconv.Atoi(p2[1])),
		})
	}

	min_x, max_x, min_y, max_y := 0, 0, 0, 0
	for _, line := range lines {
		min_x = min(min_x, line[0], line[2])
		max_x = max(max_x, line[0], line[2])
		min_y = min(min_y, line[1], line[3])
		max_y = max(max_y, line[1], line[3])
	}

	width, height := max_x-min_x+1, max_y-min_y+1
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	for _, line := range lines {
		if line[0] == line[2] {
			for y := min(line[1], line[3]); y <= max(line[1], line[3]); y++ {
				grid[y-min_y][line[0]-min_x]++
			}
		} else if line[1] == line[3] {
			for x := min(line[0], line[2]); x <= max(line[0], line[2]); x++ {
				grid[line[1]-min_y][x-min_x]++
			}
		} else {
			start_left := line[0] < line[2]
			y_start, y_end := line[1], line[3]
			if !start_left {
				y_start, y_end = y_end, y_start
			}
			dy := 1
			if y_start > y_end {
				dy = -1

			}

			for x, y := min(line[0], line[2]), y_start; x <= max(line[0], line[2]); x, y = x+1, y+dy {
				grid[y-min_y][x-min_x]++
			}
		}
	}

	count := 0
	for _, line := range grid {
		for _, val := range line {
			if val >= 2 {
				count++
			}
		}
	}
	return count
}
