package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Fold struct {
	Horizontal bool
	Pos        int
}

func solution13Part1(input string) int {
	parts := strings.Split(input, "\n\n")
	dots := parseDots(parts[0])
	folds := parseFolds(parts[1])

	dots = fold(dots, folds[0])

	return dots.len()
}

func parseDots(s string) Set[Point] {
	dots := makeSet[Point](0)
	for _, line := range strings.Split(s, "\n") {
		coords := strings.Split(line, ",")
		dots.add(Point{
			unwrap(strconv.Atoi(coords[0])),
			unwrap(strconv.Atoi(coords[1])),
		})
	}
	return dots
}

func parseFolds(s string) []Fold {
	folds := make([]Fold, 0)
	fold_regex := regexp.MustCompile("fold along (x|y)=(\\d+)")
	for _, fold := range strings.Split(s, "\n") {
		submatches := fold_regex.FindSubmatch([]byte(fold))

		horizontal := string(submatches[1]) == "y"
		pos := unwrap(strconv.Atoi(string(submatches[2])))
		folds = append(folds, Fold{horizontal, pos})
	}

	return folds
}

func fold(dots Set[Point], fold Fold) Set[Point] {
	var still_dots, moved_dots []Point
	if fold.Horizontal {
		still_dots = filter(dots.toArray(), func(p Point) bool {
			return p.Y < fold.Pos
		})
		moved_dots = filter(dots.toArray(), func(p Point) bool {
			return p.Y > fold.Pos
		})
	} else {
		still_dots = filter(dots.toArray(), func(p Point) bool {
			return p.X < fold.Pos
		})
		moved_dots = filter(dots.toArray(), func(p Point) bool {
			return p.X > fold.Pos
		})
	}

	for _, dot := range moved_dots {
		if fold.Horizontal {
			still_dots = append(still_dots, Point{dot.X, 2*fold.Pos - dot.Y})
		} else {
			still_dots = append(still_dots, Point{2*fold.Pos - dot.X, dot.Y})
		}
	}

	return makeSetFrom(still_dots)
}

func solution13Part2(input string) string {
	parts := strings.Split(input, "\n\n")
	dots := parseDots(parts[0])
	for _, f := range parseFolds(parts[1]) {
		dots = fold(dots, f)
	}

	var min_x, max_x, min_y, max_y int
	for p := range dots.data {
		min_x = min(min_x, p.X)
		max_x = max(max_x, p.X)
		min_y = min(min_y, p.Y)
		max_y = max(max_y, p.Y)
	}
	width, height := max_x-min_x+1, max_y-min_y+1

	draw := make([][]rune, 0)
	for i := 0; i < height; i++ {
		draw = append(draw, []rune(strings.Repeat(".", width)))
	}
	for dot := range dots.data {
		draw[dot.Y-min_y][dot.X-min_x] = '#'
	}

	output := ""
	for _, line := range draw {
		output += string(line) + "\n"
	}

	return output
}
