package main

import (
	"strconv"
	"strings"
)

func parse_input_17(input string) (left, top, right, bottom int) {
	input = input[13:]
	parts := apply(strings.Split(input, ", "), func(value string) []int {
		value = strings.Split(value, "=")[1]
		return apply(strings.Split(value, ".."), func(value string) int {
			return unwrap(strconv.Atoi(value))
		})
	})
	return parts[0][0], parts[1][1], parts[0][1], parts[1][0]
}

func sum_1_to_n(n int) int {
	return n * (n + 1) / 2
}

func get_x(x_vel int, t int) int {
	end_vel := max(0, x_vel-t)
	return sum_1_to_n(x_vel) - sum_1_to_n(end_vel)
}

func solution_17_1(input string) int {
	left, top, right, bottom := parse_input_17(input)

	for y_vel := -bottom - 1; y_vel >= 0; y_vel-- {
		orig_y_vel := y_vel
		// Skip past Λ arc to the second point where y == 0
		t := y_vel*2 + 1
		y_vel = -y_vel - 1

		y := 0
		valid_ts := []int{}
		for y >= bottom {
			if y <= top {
				valid_ts = append(valid_ts, t)
			}

			y += y_vel
			y_vel -= 1
			t += 1
		}

		for _, t := range valid_ts {
			for i := 0; i <= right; i++ {
				if x := get_x(i, t); x >= left && x <= right {
					return sum_1_to_n(orig_y_vel)
				}
			}
		}
	}

	panic("Target unreachable")
}

func solution_17_2(input string) int {
	left, top, right, bottom := parse_input_17(input)

	valid_count := 0
	counted := makeSet[[2]int](0)
	for orig_y_vel := abs_int(bottom) - 1; orig_y_vel >= bottom; orig_y_vel-- {
		// Skip past Λ arc to the second point where y == 0
		t := 0
		y_vel := orig_y_vel

		y := 0
		valid_ts := []int{}
		for y >= bottom {
			if y <= top {
				valid_ts = append(valid_ts, t)
			}

			y += y_vel
			y_vel -= 1
			t += 1
		}

		for _, t := range valid_ts {
			for x_vel := 0; x_vel <= right; x_vel++ {
				if counted.contains([2]int{x_vel, orig_y_vel}) {
					continue
				}

				if x := get_x(x_vel, t); x >= left && x <= right {
					counted.add([2]int{x_vel, orig_y_vel})
					valid_count += 1
				}
			}
		}
	}

	return valid_count
}
