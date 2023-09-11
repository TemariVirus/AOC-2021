package main

import (
	"math"
	"strconv"
	"strings"
)

func solution_7_1(input string) int {
	crabs := apply(strings.Split(input, ","), func(s string) int {
		return unwrap(strconv.Atoi(s))
	})

	dest := quad_curve_binary_search(crabs, func(crabs []int, idx int) int {
		left := count(crabs, func(value int) bool { return value <= idx })
		right := count(crabs, func(value int) bool { return value >= idx })
		return int(math.Abs(float64(left - right)))
	})

	return aggregate(crabs, 0, func(agg int, value int, _ int) int {
		return agg + int(math.Abs(float64(value-dest)))
	})
}

func solution_7_2(input string) int {
	crabs := apply(strings.Split(input, ","), func(s string) int {
		return unwrap(strconv.Atoi(s))
	})

	dest := quad_curve_binary_search(crabs, func(crabs []int, dest int) int {
		cost := 0
		for _, crab := range crabs {
			dist := int(math.Abs(float64(crab - dest)))
			cost += dist * (dist + 1) / 2
		}
		return cost
	})

	return aggregate(crabs, 0, func(agg int, value int, _ int) int {
		dist := int(math.Abs(float64(value - dest)))
		return agg + dist*(dist+1)/2
	})
}

func quad_curve_binary_search(arr []int, cost_func func(crabs []int, dest int) int) int {
	left := arr[0]
	right := arr[0]
	for _, crab := range arr {
		left = min(left, crab)
		right = max(right, crab)
	}

	// Binary(?) search over a quadratic curve
	for right-left > 3 {
		mid1 := (left*3 + right) / 4
		mid2 := (left*2 + right*2) / 4
		mid3 := (left + right*3) / 4
		mid1_cost := cost_func(arr, mid1)
		mid2_cost := cost_func(arr, mid2)
		mid3_cost := cost_func(arr, mid3)

		if mid1_cost < mid2_cost && mid1_cost < mid3_cost {
			right = mid2
		} else if mid2_cost < mid1_cost && mid2_cost < mid3_cost {
			left = mid1
			right = mid3
		} else {
			left = mid2
		}
	}

	return (left + right) / 2
}
