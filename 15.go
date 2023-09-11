package main

import (
	"strings"
)

func parse_input_15(input string) [][]int {
	result := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		line_risks := []int{}
		for _, char := range line {
			line_risks = append(line_risks, int(char)-int('0'))
		}
		result = append(result, line_risks)
	}
	return result
}

// A* algorithm
func least_risk_path(
	start Point,
	target Point,
	risk [][]int,
) []Point {
	open_set := []Point{start}
	closed_set := makeSet[Point](0)
	came_froms := map[Point]Point{}
	g_scores := map[Point]int{start: 0}
	f_scores := map[Point]int{start: g_scores[start] + taxicab_distance(start, target)}
	for len(open_set) > 0 {
		curr_i := aggregate(open_set, 0, func(agg int, value Point, index int) int {
			if f_scores[open_set[index]] < f_scores[open_set[agg]] {
				return index
			}
			return agg
		})
		curr := open_set[curr_i]
		open_set = append(open_set[:curr_i], open_set[curr_i+1:]...)

		if curr == target {
			// Retrace path
			path := []Point{curr}
			for curr != start {
				curr = came_froms[curr]
				path = append(path, curr)
			}
			// Reverse path
			for i := 0; i < len(path)/2; i++ {
				path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
			}
			return path
		}

		closed_set.add(curr)
		for _, neighbor := range []Point{
			{curr.X - 1, curr.Y},
			{curr.X + 1, curr.Y},
			{curr.X, curr.Y - 1},
			{curr.X, curr.Y + 1},
		} {
			if neighbor.X < 0 ||
				neighbor.X >= len(risk[0]) ||
				neighbor.Y < 0 ||
				neighbor.Y >= len(risk) ||
				closed_set.contains(neighbor) {
				continue
			}

			tentative_g_score := g_scores[curr] + risk[neighbor.Y][neighbor.X]
			if find(open_set, neighbor) == -1 {
				open_set = append(open_set, neighbor)
			}
			if g_score, ok := g_scores[neighbor]; !ok || tentative_g_score < g_score {
				came_froms[neighbor] = curr
				g_scores[neighbor] = tentative_g_score
				f_scores[neighbor] = g_scores[neighbor] + taxicab_distance(neighbor, target)
			}
		}
	}

	return []Point{}
}

func solution_15_1(input string) int {
	risks := parse_input_15(input)
	path := least_risk_path(Point{0, 0}, Point{len(risks[0]) - 1, len(risks) - 1}, risks)
	return aggregate(path[1:], 0, func(agg int, value Point, _ int) int {
		return agg + risks[value.Y][value.X]
	})
}

func solution_15_2(input string) int {
	risks_small := parse_input_15(input)

	risks := make([][]int, len(risks_small)*5)
	for i := 0; i < len(risks_small)*5; i++ {
		risks[i] = make([]int, len(risks_small[0])*5)
		for j := 0; j < len(risks_small[0])*5; j++ {
			dist := i/len(risks_small) + j/len(risks_small[0])
			x, y := j%len(risks_small[0]), i%len(risks_small)

			risk := risks_small[y][x]
			new_risk := (risk+dist-1)%9 + 1
			risks[i][j] = new_risk
		}
	}

	path := least_risk_path(Point{0, 0}, Point{len(risks[0]) - 1, len(risks) - 1}, risks)
	return aggregate(path[1:], 0, func(agg int, value Point, _ int) int {
		return agg + risks[value.Y][value.X]
	})
}
